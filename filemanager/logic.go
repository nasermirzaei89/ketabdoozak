package filemanager

import (
	"context"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/blob"
	"io"
	"log/slog"
	"path"
	"strconv"
	"strings"
	"time"
)

const DefaultExpirationDuration = time.Hour * 24

type BaseService struct {
	filesBucket *blob.Bucket
	fileRepo    FileRepository
	logger      *slog.Logger
	tracer      trace.Tracer
}

var _ Service = (*BaseService)(nil)

func NewService(filesBucket *blob.Bucket, fileRepo FileRepository) *BaseService {
	return &BaseService{
		filesBucket: filesBucket,
		fileRepo:    fileRepo,
		logger:      defaultLogger,
		tracer:      defaultTracer,
	}
}

func (svc *BaseService) putFile(ctx context.Context, filename string, reader io.Reader) error {
	w, err := svc.filesBucket.NewWriter(ctx, filename, nil)
	if err != nil {
		return errors.Wrap(err, "failed to create file writer")
	}

	defer func() {
		err := w.Close()
		if err != nil {
			svc.logger.Warn("failed to close file writer", slog.String("filename", filename), slog.String("error", err.Error()))
		}
	}()

	_, err = w.ReadFrom(reader)
	if err != nil {
		return errors.Wrap(err, "failed to read file")
	}

	return nil
}

func (svc *BaseService) uniqueFilename(ctx context.Context, filename string) (string, error) {
	filename = strings.TrimSpace(strings.ToLower(filename))

	var uniqueFilename string

	for i := 1; ; i++ {
		if i == 1 {
			uniqueFilename = filename
		} else {
			ext := path.Ext(filename)
			uniqueFilename = filename[0:len(filename)-len(ext)] + "-" + strconv.Itoa(i) + ext
		}

		exists, err := svc.filesBucket.Exists(ctx, uniqueFilename)
		if err != nil {
			return "", errors.Wrap(err, "failed to check if file exists")
		}

		if !exists {
			break
		}
	}

	return uniqueFilename, nil
}

func (svc *BaseService) UploadFile(ctx context.Context, filename string, reader io.Reader) (*File, error) {
	filename, err := svc.uniqueFilename(ctx, filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate unique filename")
	}

	err = svc.putFile(ctx, filename, reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to put file into bucket")
	}

	attributes, err := svc.filesBucket.Attributes(ctx, filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get file attributes")
	}

	createdAt := attributes.CreateTime
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	file := &File{
		Filename:    filename,
		Size:        attributes.Size,
		ContentType: attributes.ContentType,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}

	err = svc.fileRepo.Insert(ctx, file)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert file into the repo")
	}

	return file, nil
}

func (svc *BaseService) GetPublicFileURL(ctx context.Context, filename string) (string, error) {
	_, err := svc.fileRepo.Get(ctx, filename)
	if err != nil {
		return "", errors.Wrap(err, "failed to get file from repo")
	}

	signedURL, err := svc.filesBucket.SignedURL(ctx, filename, &blob.SignedURLOptions{
		Expiry:                   DefaultExpirationDuration,
		Method:                   "",
		ContentType:              "", // must be empty
		EnforceAbsentContentType: false,
		BeforeSign:               nil,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to get file signed url")
	}

	return signedURL, nil
}

func (svc *BaseService) DeleteFile(ctx context.Context, filename string) error {
	file, err := svc.fileRepo.Get(ctx, filename)
	if err != nil {
		return errors.Wrap(err, "failed to get file from repo")
	}

	err = svc.fileRepo.Delete(ctx, filename)
	if err != nil {
		return errors.Wrap(err, "failed to delete file from repo")
	}

	err = svc.filesBucket.Delete(ctx, file.Filename)
	if err != nil {
		return errors.Wrap(err, "failed to delete file from files bucket")
	}

	return nil
}
