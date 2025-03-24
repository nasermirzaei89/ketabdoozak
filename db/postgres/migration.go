package postgres

import (
	"context"
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // init file source
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
)

//go:embed migrations/*.sql
var fs embed.FS

func newMigrateInstance(sqlDB *sql.DB) (*migrate.Migrate, error) {
	sourceDriver, err := iofs.New(fs, "migrations") // Get migrations from sql folder
	if err != nil {
		return nil, errors.Wrap(err, "failed to create source driver from migrations folder")
	}

	dbDriver, err := postgres.WithInstance(sqlDB, new(postgres.Config))
	if err != nil {
		return nil, errors.Wrap(err, "could not create database driver")
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		sourceDriver,
		"postgres",
		dbDriver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not create migration instance")
	}

	return m, nil
}

func MigrateUp(ctx context.Context, sqlDB *sql.DB) error {
	m, err := newMigrateInstance(sqlDB)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "could not run migrations")
	}

	logger.InfoContext(ctx, "Migration applied successfully.")

	return nil
}

func MigrateDown(ctx context.Context, sqlDB *sql.DB) error {
	m, err := newMigrateInstance(sqlDB)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	err = m.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "could not migrate down")
	}

	logger.InfoContext(ctx, "Migration down applied successfully.")

	return nil
}
