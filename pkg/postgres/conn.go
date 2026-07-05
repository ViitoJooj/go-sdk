// Package postgres exposes a shared *sql.DB opened once per process and a
// convenience wrapper around golang-migrate for schema migrations.
package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is the shared connection pool set by Conn. Callers dereference it directly
// (mirrors the auth service's original API).
var DB *sql.DB

// Conn opens a lib/pq connection to dsn, pings it, and stores the pool in DB.
// Returns an error instead of calling log.Fatal so services can decide how to
// fail (retry, panic, exit with a specific code) at startup.
func Conn(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open postgres: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping postgres: %w", err)
	}
	DB = db
	return nil
}
