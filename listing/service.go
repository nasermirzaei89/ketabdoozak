package listing

import (
	"github.com/go-playground/validator/v10"
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
