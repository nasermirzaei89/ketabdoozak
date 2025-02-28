package listing

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
)

type Service struct {
	itemRepo ItemRepository
	validate *validator.Validate
	logger   *slog.Logger
	tracer   trace.Tracer
}

func NewService(itemRepo ItemRepository, validate *validator.Validate) *Service {
	return &Service{
		itemRepo: itemRepo,
		validate: validate,
		logger:   defaultLogger,
		tracer:   defaultTracer,
	}
}

type ListItemsRequest struct {
	Query   string
	OwnerID string
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

	return item, nil
}

func (svc *Service) DeleteItem(ctx context.Context, itemID string) error {
	ctx, span := svc.tracer.Start(ctx, "DeleteItem")
	defer span.End()

	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	err = svc.itemRepo.Delete(ctx, item.ID)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())

		return errors.Wrapf(err, "error on delete item with id '%s'", itemID)
	}

	return nil
}
