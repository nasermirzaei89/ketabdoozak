package postgres

import (
	"context"
	"database/sql"
	goerrors "errors"
	"github.com/XSAM/otelsql"
	_ "github.com/lib/pq" // init postgres driver
	"github.com/pkg/errors"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

func NewDB(ctx context.Context, dbDSN string) (*sql.DB, func() error, error) {
	sqlDB, err := otelsql.Open("postgres", dbDSN, otelsql.WithAttributes(semconv.DBSystemPostgreSQL))
	if err != nil {
		return nil, nil, errors.Wrap(err, "error opening db")
	}

	closeFunc := func() error {
		err := sqlDB.Close()
		if err != nil {
			return errors.Wrap(err, "error closing db")
		}

		return nil
	}

	err = otelsql.RegisterDBStatsMetrics(sqlDB, otelsql.WithAttributes(semconv.DBSystemPostgreSQL))
	if err != nil {
		err = goerrors.Join(err, closeFunc())

		return nil, nil, errors.Wrap(err, "error registering stats metrics")
	}

	err = sqlDB.Ping()
	if err != nil {
		err = goerrors.Join(err, closeFunc())

		return nil, nil, errors.Wrap(err, "error pinging db")
	}

	err = MigrateUp(ctx, sqlDB)
	if err != nil {
		err = goerrors.Join(err, closeFunc())

		return nil, nil, errors.Wrap(err, "error migrating db")
	}

	return sqlDB, closeFunc, nil
}
