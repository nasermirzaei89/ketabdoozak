package listing

import "fmt"

type LocationWithIDNotFoundError struct {
	ID string
}

func (err LocationWithIDNotFoundError) Error() string {
	return fmt.Sprintf("location with id '%s' not found", err.ID)
}

type ItemWithIDNotFoundError struct {
	ID string
}

func (err ItemWithIDNotFoundError) Error() string {
	return fmt.Sprintf("item with id '%s' not found", err.ID)
}

type CannotSendItemForReviewError struct {
	ID     string
	Status ItemStatus
}

func (err CannotSendItemForReviewError) Error() string {
	return fmt.Sprintf("cannot send item with id '%s' and status '%s' for review", err.ID, err.Status)
}

type CannotPublishItemError struct {
	ID     string
	Status ItemStatus
}

func (err CannotPublishItemError) Error() string {
	return fmt.Sprintf("cannot publish item with id '%s' status '%s'", err.ID, err.Status)
}

type CannotArchiveItemError struct {
	ID     string
	Status ItemStatus
}

func (err CannotArchiveItemError) Error() string {
	return fmt.Sprintf("cannot archive item with id '%s' status '%s'", err.ID, err.Status)
}

type CannotDeleteItemError struct {
	ID     string
	Status ItemStatus
}

func (err CannotDeleteItemError) Error() string {
	return fmt.Sprintf("cannot delete item with id '%s' status '%s'", err.ID, err.Status)
}
