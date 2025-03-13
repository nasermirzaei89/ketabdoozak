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

func (svc *BaseService) UploadFile(ctx context.Context, filename string, reader io.Reader, contentType string) (*File, error) {
	filename = strings.ToLower(filename)

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
			return nil, errors.Wrap(err, "failed to check if file exists")
		}

		if !exists {
			break
		}
	}

	w, err := svc.filesBucket.NewWriter(ctx, uniqueFilename, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file writer")
	}

	defer func() {
		err := w.Close()
		if err != nil {
			svc.logger.Warn("failed to close file writer", slog.String("filename", uniqueFilename), slog.String("error", err.Error()))
		}
	}()

	fileSize, err := w.ReadFrom(reader)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	file := &File{
		Filename:    uniqueFilename,
		Size:        fileSize,
		ContentType: contentType,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
