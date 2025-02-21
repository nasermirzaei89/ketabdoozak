package postgres

import (
	"go.opentelemetry.io/contrib/bridges/otelslog"
)

const packageName = "github.com/nasermirzaei89/ketabdoozak/db/postgres"

var logger = otelslog.NewLogger(packageName) //nolint:gochecknoglobals
