package filemanager

import (
	"fmt"
)

type FileByFilenameNotFoundError struct {
	FileName string
}

func (err FileByFilenameNotFoundError) Error() string {
	return fmt.Sprintf("file with filename %s not found", err.FileName)
}
