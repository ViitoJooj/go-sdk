// Package postgres exposes a shared *sql.DB opened once per process and a
// convenience wrapper around golang-migrate for schema migrations.
package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// DB is the shared connection pool set by Conn.
var DB *sql.DB

// Config configures the shared connection pool.
type Config struct {
	DSN             string
	MaxOpenConns    int           // 0 = driver default (unlimited)
	MaxIdleConns    int           // 0 = driver default (2)
	ConnMaxLifetime time.Duration // 0 = unlimited
}

// Conn opens a lib/pq connection to cfg.DSN, pings it, and stores the pool in
// DB. Returns an error instead of calling log.Fatal so services can decide how
// to fail at startup.
func Conn(cfg Config) error {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return fmt.Errorf("open postgres: %w", err)
	}
	if cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping postgres: %w", err)
	}
	DB = db
	return nil
}
