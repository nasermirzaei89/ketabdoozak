package filemanager

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log/slog"
	"path"
	"strconv"
	"strings"
	"time"
)

const DefaultExpirationDuration = time.Hour * 24

type BaseService struct {
	minioClient     *minio.Client
	minioBucketName string
	fileRepo        FileRepository
	logger          *slog.Logger
	tracer          trace.Tracer
}

var _ Service = (*BaseService)(nil)

func NewService(minioClient *minio.Client, minioBucketName string, fileRepo FileRepository) *BaseService {
	return &BaseService{
		minioClient:     minioClient,
		minioBucketName: minioBucketName,
		fileRepo:        fileRepo,
		logger:          defaultLogger,
		tracer:          defaultTracer,
	}
}

func (svc *BaseService) UploadFile(ctx context.Context, filename string, reader io.Reader, fileSize int64, contentType string) (*File, error) {
	filename = strings.ToLower(filename)

	var uniqueFilename string

	for i := 1; ; i++ {
		if i == 1 {
			uniqueFilename = filename
		} else {
			ext := path.Ext(filename)
			uniqueFilename = filename[0:len(filename)-len(ext)] + "-" + strconv.Itoa(i) + ext
		}

		_, err := svc.minioClient.StatObject(ctx, svc.minioBucketName, uniqueFilename, minio.GetObjectOptions{})
		if err != nil {
			var minioErr minio.ErrorResponse
			if errors.As(err, &minioErr) && minioErr.Code == "NoSuchKey" {
				break
			}

			return nil, errors.Wrap(err, "failed to check if file exists")
		}
	}

	info, err := svc.minioClient.PutObject(ctx, svc.minioBucketName, uniqueFilename, reader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to put object to minio storage")
	}

	file := &File{
		Filename:    uniqueFilename,
		Size:        info.Size,
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
	file, err := svc.fileRepo.Get(ctx, filename)
	if err != nil {
		return "", errors.Wrap(err, "failed to get file from repo")
	}

	u, err := svc.minioClient.PresignedGetObject(ctx, svc.minioBucketName, file.Filename, DefaultExpirationDuration, nil)
	if err != nil {
		return "", errors.Wrap(err, "failed to presigned get object from minio")
	}

	return u.String(), nil
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

	opts := minio.RemoveObjectOptions{
		ForceDelete:      false,
		GovernanceBypass: false,
		VersionID:        "",
		Internal:         minio.AdvancedRemoveOptions{},
	}

	err = svc.minioClient.RemoveObject(ctx, svc.minioBucketName, file.Filename, opts)
	if err != nil {
		return errors.Wrap(err, "failed to presigned get object from minio")
	}

	return nil
}
