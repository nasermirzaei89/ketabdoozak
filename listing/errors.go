package listing

import "fmt"

type ItemWithIDNotFoundError struct {
	ID string
}

func (err ItemWithIDNotFoundError) Error() string {
	return fmt.Sprintf("item with id '%s' not found", err.ID)
}
