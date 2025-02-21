package postgres

//
//import (
//	"context"
//	"database/sql"
//	"github.com/Masterminds/squirrel"
//	"github.com/nasermirzaei89/ketabdoozak/listing"
//	"github.com/pkg/errors"
//	"log"
//)
//
//type ListingItemRepo struct {
//	db    *sql.DB
//	table string
//	cols  []string
//}
//
//func (repo *ListingItemRepo) Get(ctx context.Context, itemID string) (*listing.Item, error) {
//	q := squirrel.Select(repo.cols...).From(repo.table).Where(squirrel.Eq{"id": itemID})
//
//	var item listing.Item
//
//	err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).QueryRowContext(ctx).Scan(
//		&item.ID,
//		&item.Title,
//		&item.Excerpt,
//		&item.Description,
//		&item.ThumbnailURL,
//		&item.InstructorID,
//		&item.Status,
//	)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, listing.ItemWithIDNotFoundError{ID: itemID}
//		}
//
//		return nil, errors.Wrap(err, "query row failed")
//	}
//
//	return &item, nil
//}
//
//func (repo *ListingItemRepo) List(ctx context.Context, req listing.ListItemsRequest) ([]listing.Item, error) {
//	q := squirrel.Select(repo.cols...).From(repo.table)
//
//	rows, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).QueryContext(ctx)
//	if err != nil {
//		return nil, errors.Wrap(err, "query failed")
//	}
//
//	defer func() {
//		err := rows.Close()
//		if err != nil {
//			log.Printf("failed to close rows: %v", err)
//		}
//	}()
//
//	itemList := make([]listing.Item, 0)
//
//	for rows.Next() {
//		var item listing.Item
//
//		err := rows.Scan(
//			&item.ID,
//			&item.Title,
//			&item.Excerpt,
//			&item.Description,
//			&item.ThumbnailURL,
//			&item.InstructorID,
//			&item.Status,
//		)
//		if err != nil {
//			return nil, errors.Wrap(err, "scan row failed")
//		}
//
//		itemList = append(itemList, item)
//	}
//
//	err = rows.Close()
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to close rows")
//	}
//
//	err = rows.Err()
//	if err != nil {
//		return nil, errors.Wrap(err, "error encountered during rows iteration")
//	}
//
//	return itemList, nil
//}
//
//func (repo *ListingItemRepo) Insert(ctx context.Context, item *listing.Item) error {
//	q := squirrel.Insert(repo.table).Columns(repo.cols...).Values(
//		item.ID,
//		item.Title,
//		item.Excerpt,
//		item.Description,
//		item.ThumbnailURL,
//		item.InstructorID,
//		item.Status,
//	)
//
//	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
//	if err != nil {
//		return errors.Wrap(err, "execute query failed")
//	}
//
//	return nil
//}
//
//func (repo *ListingItemRepo) Replace(ctx context.Context, itemID string, item *listing.Item) error {
//	q := squirrel.Update(repo.table).SetMap(map[string]any{
//		"id":            item.ID,
//		"title":         item.Title,
//		"excerpt":       item.Excerpt,
//		"description":   item.Description,
//		"thumbnail_url": item.ThumbnailURL,
//		"instructor_id": item.InstructorID,
//		"status":        item.Status,
//	}).Where(squirrel.Eq{"id": itemID})
//
//	_, err := q.RunWith(repo.db).PlaceholderFormat(squirrel.Dollar).ExecContext(ctx)
//	if err != nil {
//		return errors.Wrap(err, "execute query failed")
//	}
//
//	return nil
//}
//
//var _ listing.ItemRepository = (*ListingItemRepo)(nil)
//
//func NewListingItemRepo(db *sql.DB) *ListingItemRepo {
//	return &ListingItemRepo{
//		db:    db,
//		table: "listing_items",
//		cols: []string{
//			"id",
//			"title",
//			"excerpt",
//			"description",
//			"thumbnail_url",
//			"instructor_id",
//			"status",
//		},
//	}
//}
