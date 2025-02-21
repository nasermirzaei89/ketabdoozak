package listing

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

const name = "github.com/nasermirzaei89/ketabdoozak/listing"

var (
	defaultLogger = otelslog.NewLogger(name)
	defaultTracer = otel.Tracer(name)
)
