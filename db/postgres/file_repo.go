package postgres

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/nasermirzaei89/ketabdoozak/filemanager"
	"github.com/pkg/errors"
)

type FileRepo struct {
	db    *sql.DB
	table string
	cols  []string
}

func (repo *FileRepo) Insert(ctx context.Context, file *filemanager.File) error {
	q := squirrel.Insert(repo.table).Columns(repo.cols...).Values(
		file.Filename,
		file.Size,
		file.ContentType,
		file.CreatedAt,
		file.UpdatedAt,
	)

	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "execute query failed")
	}

	return nil
}

func (repo *FileRepo) Get(ctx context.Context, filename string) (*filemanager.File, error) {
	q := squirrel.Select(repo.cols...).From(repo.table).Where(squirrel.Eq{"filename": filename})

	var file filemanager.File

	err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).QueryRowContext(ctx).Scan(
		&file.Filename,
		&file.Size,
		&file.ContentType,
		&file.CreatedAt,
		&file.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, filemanager.FileByFilenameNotFoundError{FileName: filename}
		}

		return nil, errors.Wrap(err, "query row failed")
	}

	return &file, nil
}

func (repo *FileRepo) Delete(ctx context.Context, filename string) error {
	q := squirrel.Delete(repo.table).Where(squirrel.Eq{"filename": filename})

	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "query row failed")
	}

	return nil
}

var _ filemanager.FileRepository = (*FileRepo)(nil)

func NewFileRepo(db *sql.DB) *FileRepo {
	return &FileRepo{
		db:    db,
		table: "file_manager_files",
		cols: []string{
			"filename",
			"size",
			"content_type",
			"created_at",
			"updated_at",
		},
	}
}
