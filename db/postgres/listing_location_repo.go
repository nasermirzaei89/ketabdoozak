package postgres

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/pkg/errors"
	"log"
)

type ListingLocationRepo struct {
	db    *sql.DB
	table string
	cols  []string
}

func (repo *ListingLocationRepo) scanRow(rowScanner squirrel.RowScanner) (*listing.Location, error) {
	var location listing.Location

	err := rowScanner.Scan(
		&location.ID,
		&location.Title,
		&location.ParentID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "scan row failed")
	}

	return &location, nil
}

func (repo *ListingLocationRepo) Get(ctx context.Context, locationID string) (*listing.Location, error) {
	q := squirrel.Select(repo.cols...).From(repo.table).Where(squirrel.Eq{"id": locationID})

	row := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).QueryRowContext(ctx)

	location, err := repo.scanRow(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, listing.LocationWithIDNotFoundError{ID: locationID}
		}

		return nil, errors.Wrap(err, "query row failed")
	}

	return location, nil
}

func (repo *ListingLocationRepo) List(ctx context.Context) ([]*listing.Location, error) {
	q := squirrel.Select(repo.cols...).From(repo.table)

	rows, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).QueryContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	locationList := make([]*listing.Location, 0)

	for rows.Next() {
		location, err := repo.scanRow(rows)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}

		locationList = append(locationList, location)
	}

	err = rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close rows")
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error encountered during rows iteration")
	}

	return locationList, nil
}

func (repo *ListingLocationRepo) Insert(ctx context.Context, location *listing.Location) error {
	q := squirrel.Insert(repo.table).Columns(repo.cols...).Values(
		location.ID,
		location.Title,
		location.ParentID,
	)

	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "execute query failed")
	}

	return nil
}

func (repo *ListingLocationRepo) Replace(ctx context.Context, locationID string, location *listing.Location) error {
	q := squirrel.Update(repo.table).SetMap(map[string]any{
		"id":        location.ID,
		"title":     location.Title,
		"parent_id": location.ParentID,
	}).Where(squirrel.Eq{"id": locationID})

	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "execute query failed")
	}

	return nil
}

var _ listing.LocationRepository = (*ListingLocationRepo)(nil)

func NewListingLocationRepo(db *sql.DB) *ListingLocationRepo {
	return &ListingLocationRepo{
		db:    db,
		table: "listing_locations",
		cols: []string{
			"id",
			"title",
			"parent_id",
		},
	}
}
