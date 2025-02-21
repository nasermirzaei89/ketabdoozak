package authorization

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

const name = "github.com/nasermirzaei89/ketabdoozak/authorization"

var (
	logger = otelslog.NewLogger(name)
	tracer = otel.Tracer(name)
)
