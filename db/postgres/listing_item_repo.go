package postgres

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/nasermirzaei89/extypes"
	"github.com/nasermirzaei89/ketabdoozak/listing"
	"github.com/pkg/errors"
	"log"
)

type ListingItemRepo struct {
	db    *sql.DB
	table string
	cols  []string
}

func (repo *ListingItemRepo) scanRow(rowScanner squirrel.RowScanner) (*listing.Item, error) {
	var item listing.Item

	var (
		itemTypes       extypes.JSONObject
		itemContactInfo extypes.JSONObject
	)

	err := rowScanner.Scan(
		&item.ID,
		&item.Title,
		&item.OwnerID,
		&item.OwnerName,
		&item.LocationID,
		&item.LocationTitle,
		&itemTypes,
		&itemContactInfo,
		&item.Description,
		&item.Status,
		&item.Lent,
		&item.ThumbnailURL,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.PublishedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "scan row failed")
	}

	err = itemTypes.Decode(&item.Types)
	if err != nil {
		return nil, errors.Wrap(err, "decode itemTypes failed")
	}

	err = itemContactInfo.Decode(&item.ContactInfo)
	if err != nil {
		return nil, errors.Wrap(err, "contact info decoding failed")
	}

	return &item, nil
}

func (repo *ListingItemRepo) Get(ctx context.Context, itemID string) (*listing.Item, error) {
	q := squirrel.Select(repo.cols...).From(repo.table).Where(squirrel.Eq{"id": itemID})

	row := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).QueryRowContext(ctx)

	item, err := repo.scanRow(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, listing.ItemWithIDNotFoundError{ID: itemID}
		}

		return nil, errors.Wrap(err, "query row failed")
	}

	return item, nil
}

func (repo *ListingItemRepo) List(ctx context.Context, req *listing.ListItemsRequest) ([]*listing.Item, error) {
	q := squirrel.Select(repo.cols...).From(repo.table)

	if req != nil {
		if req.Query != "" {
			q = q.Where("(to_tsvector(title) || location_title || description) @@ plainto_tsquery(?)", req.Query)
		}

		if req.OwnerID != "" {
			q = q.Where(squirrel.Eq{"owner_id": req.OwnerID})
		}

		if req.Status != "" {
			q = q.Where(squirrel.Eq{"status": req.Status})
		}
	}

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

	itemList := make([]*listing.Item, 0)

	for rows.Next() {
		item, err := repo.scanRow(rows)
		if err != nil {
			return nil, errors.Wrap(err, "scan row failed")
		}

		itemList = append(itemList, item)
	}

	err = rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to close rows")
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.Wrap(err, "error encountered during rows iteration")
	}

	return itemList, nil
}

func (repo *ListingItemRepo) Insert(ctx context.Context, item *listing.Item) error {
	itemTypes := extypes.NewJSONObject(item.Types)
	itemContactInfo := extypes.NewJSONObject(item.ContactInfo)
	q := squirrel.Insert(repo.table).Columns(repo.cols...).Values(
		item.ID,
		item.Title,
		item.OwnerID,
		item.OwnerName,
		item.LocationID,
		item.LocationTitle,
		itemTypes,
		itemContactInfo,
		item.Description,
		item.Status,
		item.Lent,
		item.ThumbnailURL,
		item.CreatedAt,
		item.UpdatedAt,
		item.PublishedAt,
	)

	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "execute query failed")
	}

	return nil
}

func (repo *ListingItemRepo) Replace(ctx context.Context, itemID string, item *listing.Item) error {
	itemTypes := extypes.NewJSONObject(item.Types)
	itemContactInfo := extypes.NewJSONObject(item.ContactInfo)

	q := squirrel.Update(repo.table).SetMap(map[string]any{
		"id":             item.ID,
		"title":          item.Title,
		"owner_id":       item.OwnerID,
		"owner_name":     item.OwnerName,
		"location_id":    item.LocationID,
		"location_title": item.LocationTitle,
		"types":          itemTypes,
		"contact_info":   itemContactInfo,
		"description":    item.Description,
		"status":         item.Status,
		"lent":           item.Lent,
		"thumbnail_url":  item.ThumbnailURL,
		"created_at":     item.CreatedAt,
		"updated_at":     item.UpdatedAt,
		"published_at":   item.PublishedAt,
	}).Where(squirrel.Eq{"id": itemID})

	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, "execute query failed")
	}

	return nil
}

var _ listing.ItemRepository = (*ListingItemRepo)(nil)

func NewListingItemRepo(db *sql.DB) *ListingItemRepo {
	return &ListingItemRepo{
		db:    db,
		table: "listing_items",
		cols: []string{
			"id",
			"title",
			"owner_id",
			"owner_name",
			"location_id",
			"location_title",
			"types",
			"contact_info",
			"description",
			"status",
			"lent",
			"thumbnail_url",
			"created_at",
			"updated_at",
			"published_at",
		},
	}
}
