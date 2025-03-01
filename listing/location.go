package listing

import "context"

type Location struct {
	ID       string
	Title    string
	ParentID string
}

type LocationRepository interface {
	List(ctx context.Context) (locations []*Location, err error)
	Get(ctx context.Context, locationID string) (location *Location, err error)
}
