package listing

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

var (
	defaultLogger = otelslog.NewLogger(ServiceName)
	defaultTracer = otel.Tracer(ServiceName)
)
