package listing

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
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
	Query string
}

type ListItemsResponse struct {
	Items []*Item `json:"items"`
}

func (svc *Service) ListItems(ctx context.Context, req *ListItemsRequest) (*ListItemsResponse, error) {
	items, err := svc.itemRepo.List(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error on list items")
	}

	rsp := &ListItemsResponse{
		Items: items,
	}

	return rsp, nil
}

func (svc *Service) GetItem(ctx context.Context, itemID string) (*Item, error) {
	item, err := svc.itemRepo.Get(ctx, itemID)
	if err != nil {
		return nil, errors.Wrapf(err, "error on get item with id '%s'", itemID)
	}

	return item, nil
}
