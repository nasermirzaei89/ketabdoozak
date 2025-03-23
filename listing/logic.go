package listing

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nasermirzaei89/ketabdoozak/sharedcontext"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"html/template"
	"log/slog"
	"slices"
	"time"
)

type BaseService struct {
	locationRepo LocationRepository
	itemRepo     ItemRepository
	validate     *validator.Validate
	htmlPolicy   *bluemonday.Policy
	logger       *slog.Logger
	tracer       trace.Tracer
}

var _ Service = (*BaseService)(nil)

func NewService(locationRepo LocationRepository, itemRepo ItemRepository, validate *validator.Validate) *BaseService {
	return &BaseService{
		locationRepo: locationRepo,
		itemRepo:     itemRepo,
		validate:     validate,
		htmlPolicy:   bluemonday.UGCPolicy(),
		logger:       defaultLogger,
		tracer:       defaultTracer,
	}
}

func (svc *BaseService) ListLocations(ctx context.Context) (*ListLocationsResponse, error) {
	locations, err := svc.locationRepo.List(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error on list locations")
	}

	rsp := &ListLocationsResponse{
		Items: locations,
	}

	return rsp, nil
}

func (svc *BaseService) ListPublishedItems(ctx context.Context, q string) (*ListItemsResponse, error) {
	ctx, span := svc.tracer.Start(ctx, "ListPublishedItems")
	defer span.End()

	req := &ListItemsRequest{
		Query:  q,
		Status: ItemStatusPublished,
	}

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

func (svc *BaseService) ListMyItems(ctx context.Context, q string) (*ListItemsResponse, error) {
	ctx, span := svc.tracer.Start(ctx, "ListMyItems")
	defer span.End()

	req := &ListItemsRequest{
		Query:   q,
		OwnerID: sharedcontext.GetSubject(ctx),
	}

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

func (svc *BaseService) GetItem(ctx context.Context, itemID string) (*Item, error) {
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

func (svc *BaseService) GetPublishedItem(ctx context.Context, itemID string) (*Item, error) {
	ctx, span := svc.tracer.Start(ctx, "GetPublishedItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	if item.Status != ItemStatusPublished {
		return nil, ItemWithIDNotFoundError{
			ID: item.ID,
		}
	}

	return item, nil
}

func (svc *BaseService) SendItemForPublish(ctx context.Context, itemID string) error {
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

func (svc *BaseService) PublishItem(ctx context.Context, itemID string) error {
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

func (svc *BaseService) ArchiveItem(ctx context.Context, itemID string) error {
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

func (svc *BaseService) DeleteItem(ctx context.Context, itemID string) error {
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

func (svc *BaseService) CreateItem(ctx context.Context, req *CreateItemRequest) (*Item, error) {
	ctx, span := svc.tracer.Start(ctx, "CreateItem")
	defer span.End()

	username := sharedcontext.GetSubject(ctx)

	location, err := svc.locationRepo.Get(ctx, req.LocationID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return nil, errors.Wrap(err, "error on get location")
	}

	req.Description = template.HTML(svc.htmlPolicy.Sanitize(string(req.Description)))

	timeNow := time.Now()

	item := &Item{
		ID:            slug.Make(req.Title),
		Title:         req.Title,
		OwnerID:       username,
		OwnerName:     req.OwnerName,
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

func (svc *BaseService) UpdateItem(ctx context.Context, itemID string, req *UpdateItemRequest) error {
	ctx, span := svc.tracer.Start(ctx, "UpdateItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	item.Title = req.Title
	item.OwnerName = req.OwnerName
	item.LocationID = req.LocationID

	location, err := svc.locationRepo.Get(ctx, req.LocationID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrap(err, "error on get location")
	}

	item.LocationTitle = location.Title

	item.Types = req.Types
	item.ContactInfo = req.ContactInfo

	req.Description = template.HTML(svc.htmlPolicy.Sanitize(string(req.Description)))

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
