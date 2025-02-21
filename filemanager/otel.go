package filemanager

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
)

const name = "github.com/nasermirzaei89/ketabdoozak/filemanager"

var (
	defaultLogger = otelslog.NewLogger(name)
	defaultTracer = otel.Tracer(name)
)
