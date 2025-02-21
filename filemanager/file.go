package filemanager

import (
	"context"
	"time"
)

type File struct {
	Filename    string    `json:"filename"`
	Size        int64     `json:"size"`
	ContentType string    `json:"contentType"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type FileRepository interface {
	Insert(ctx context.Context, file *File) (err error)
	Get(ctx context.Context, filename string) (file *File, err error)
}
