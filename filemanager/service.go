package filemanager

import (
	"context"
	"io"
)

const ServiceName = "github.com/nasermirzaei89/ketabdoozak/filemanager"

type Service interface {
	UploadFile(ctx context.Context, filename string, reader io.Reader) (file *File, err error)
	GetPublicFileURL(ctx context.Context, filename string) (fileURL string, err error)
	DeleteFile(ctx context.Context, filename string) (err error)
}
