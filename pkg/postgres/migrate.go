package postgres

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate applies all up-migrations from sourceURL against DB. sourceURL is any
// go-migrate source string, e.g. "file://migrations". Requires Conn to have run.
func Migrate(sourceURL string) error {
	if DB == nil {
		return fmt.Errorf("postgres: DB not initialized; call Conn first")
	}
	driver, err := pgmigrate.WithInstance(DB, &pgmigrate.Config{})
	if err != nil {
		return fmt.Errorf("migration driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(sourceURL, "postgres", driver)
	if err != nil {
		return fmt.Errorf("migrator init: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrations up: %w", err)
	}
	return nil
}
