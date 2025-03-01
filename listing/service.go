package listing

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"html/template"
	"log/slog"
	"slices"
	"time"
)

type Service struct {
	locationRepo LocationRepository
	itemRepo     ItemRepository
	validate     *validator.Validate
	logger       *slog.Logger
	tracer       trace.Tracer
}

func NewService(locationRepo LocationRepository, itemRepo ItemRepository, validate *validator.Validate) *Service {
	return &Service{
		locationRepo: locationRepo,
		itemRepo:     itemRepo,
		validate:     validate,
		logger:       defaultLogger,
		tracer:       defaultTracer,
	}
}

type ListLocationsResponse struct {
	Items []*Location `json:"items"`
}

func (svc *Service) ListLocations(ctx context.Context) (*ListLocationsResponse, error) {
	locations, err := svc.locationRepo.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list locations")
	}

	rsp := &ListLocationsResponse{
		Items: locations,
	}

	return rsp, nil
}

type ListItemsRequest struct {
	Query   string
	OwnerID string
	Status  ItemStatus
}

type ListItemsResponse struct {
	Items []*Item `json:"items"`
}

func (svc *Service) ListItems(ctx context.Context, req *ListItemsRequest) (*ListItemsResponse, error) {
	ctx, span := svc.tracer.Start(ctx, "ListItems")
	defer span.End()

	items, err := svc.itemRepo.List(ctx, req)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, errors.Wrap(err, "error on list items")
	}

	rsp := &ListItemsResponse{
		Items: items,
	}

	return rsp, nil
}

func (svc *Service) GetItem(ctx context.Context, itemID string) (*Item, error) {
	ctx, span := svc.tracer.Start(ctx, "GetItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	if item.OwnerID != sharedcontext.GetSubject(ctx) && item.Status != ItemStatusPublished {
		return nil, ItemWithIDNotFoundError{
			ID: item.ID,
		}
	}

	return item, nil
}

func (svc *Service) SendItemForPublish(ctx context.Context, itemID string) error {
	ctx, span := svc.tracer.Start(ctx, "SendItemForPublish")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	if item.Status != ItemStatusDraft {
		span.SetStatus(codes.Error, "item is not draft")

		return CannotSendItemForReviewError{
			ID:     item.ID,
			Status: item.Status,
		}
	}

	item.Status = ItemStatusPendingReview
	item.UpdatedAt = time.Now()

	err = svc.itemRepo.Replace(ctx, item.ID, item)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on replace item with id '%s'", itemID)
	}

	return nil
}

func (svc *Service) PublishItem(ctx context.Context, itemID string) error {
	ctx, span := svc.tracer.Start(ctx, "PublishItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	if !slices.Contains([]ItemStatus{ItemStatusExpired, ItemStatusArchived}, item.Status) {
		span.SetStatus(codes.Error, "cannot publish item")

		return CannotPublishItemError{
			ID:     item.ID,
			Status: item.Status,
		}
	}

	item.Status = ItemStatusPublished
	item.UpdatedAt = time.Now()

	err = svc.itemRepo.Replace(ctx, item.ID, item)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on replace item with id '%s'", itemID)
	}

	return nil
}

func (svc *Service) ArchiveItem(ctx context.Context, itemID string) error {
	ctx, span := svc.tracer.Start(ctx, "ArchiveItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	if !slices.Contains([]ItemStatus{ItemStatusPublished, ItemStatusExpired}, item.Status) {
		span.SetStatus(codes.Error, "cannot archive item")

		return CannotArchiveItemError{
			ID:     item.ID,
			Status: item.Status,
		}
	}

	item.Status = ItemStatusArchived
	item.UpdatedAt = time.Now()

	err = svc.itemRepo.Replace(ctx, item.ID, item)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on replace item with id '%s'", itemID)
	}

	return nil
}

func (svc *Service) DeleteItem(ctx context.Context, itemID string) error {
	ctx, span := svc.tracer.Start(ctx, "DeleteItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	if !slices.Contains([]ItemStatus{ItemStatusDraft, ItemStatusArchived}, item.Status) {
		span.SetStatus(codes.Error, "cannot delete item")

		return CannotDeleteItemError{
			ID:     item.ID,
			Status: item.Status,
		}
	}

	item.Status = ItemStatusDeleted
	item.UpdatedAt = time.Now()

	err = svc.itemRepo.Replace(ctx, item.ID, item)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on replace item with id '%s'", itemID)
	}

	return nil
}

type CreateItemRequest struct {
	Title        string
	LocationID   string
	Types        []ItemType
	ContactInfo  []ItemContactInfo
	Description  template.HTML
	ThumbnailURL string
	AsDraft      bool
}

func (svc *Service) CreateItem(ctx context.Context, req *CreateItemRequest) (*Item, error) {
	ctx, span := svc.tracer.Start(ctx, "CreateItem")
	defer span.End()

	username := sharedcontext.GetSubject(ctx)

	location, err := svc.locationRepo.Get(ctx, req.LocationID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, errors.Wrap(err, "error on get location")
	}

	timeNow := time.Now()

	item := &Item{
		ID:            slug.Make(req.Title),
		Title:         req.Title,
		OwnerID:       username,
		OwnerName:     "", // FIXME
		LocationID:    location.ID,
		LocationTitle: location.Title,
		Types:         req.Types,
		ContactInfo:   req.ContactInfo,
		Description:   req.Description,
		Status:        ItemStatusPendingReview,
		Lent:          false,
		ThumbnailURL:  req.ThumbnailURL,
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		PublishedAt:   nil,
	}

	if req.AsDraft {
		item.Status = ItemStatusDraft
	}

	err = svc.itemRepo.Insert(ctx, item)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, errors.Wrap(err, "error on insert item")
	}

	return item, nil
}

type UpdateItemRequest struct {
	Title        string
	LocationID   string
	Types        []ItemType
	ContactInfo  []ItemContactInfo
	Description  template.HTML
	Lent         bool
	ThumbnailURL string
	AsDraft      bool
}

func (svc *Service) UpdateItem(ctx context.Context, itemID string, req *UpdateItemRequest) error {
	ctx, span := svc.tracer.Start(ctx, "UpdateItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	item.Title = req.Title
	item.LocationID = req.LocationID

	location, err := svc.locationRepo.Get(ctx, req.LocationID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrap(err, "error on get location")
	}

	item.LocationTitle = location.Title

	item.Types = req.Types
	item.ContactInfo = req.ContactInfo
	item.Description = req.Description
	item.Lent = req.Lent
	item.ThumbnailURL = req.ThumbnailURL

	timeNow := time.Now()

	if req.AsDraft {
		item.Status = ItemStatusDraft
	} else {
		item.Status = ItemStatusPendingReview
	}

	item.UpdatedAt = timeNow

	err = svc.itemRepo.Replace(ctx, item.ID, item)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on replace item with id '%s'", itemID)
	}

	return nil
}
