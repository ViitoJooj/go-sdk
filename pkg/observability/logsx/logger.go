// Package logsx sets up per-stream slog loggers with rotating on-disk files.
// Streams are labelled ("app", "access", "security", "error") so downstream
// log shippers (Promtail, Vector) can route by the "logger" attribute.
package logsx

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

// Log streams. Each is a separate rotating file plus a stdout mirror so the
// container's stdout carries every line while on-disk files keep an audit copy.
const (
	streamApp      = "app"      // general application + internal errors
	streamAccess   = "access"   // one line per HTTP request
	streamSecurity = "security" // auth events: login/logout/register/oauth/token
	streamError    = "error"    // LevelError records only
)

const (
	maxLines = 500 // rotate every 500 lines
	keepFile = 5   // retained rotated files per stream
)

var (
	appLogger      = slog.Default()
	accessLogger   = slog.Default()
	securityLogger = slog.Default()
	errorLogger    = slog.Default()

	writers []*RotatingWriter // tracked for Close
)

// App returns the general application logger.
func App() *slog.Logger { return appLogger }

// Access returns the HTTP access logger.
func Access() *slog.Logger { return accessLogger }

// Security returns the auth/security audit logger.
func Security() *slog.Logger { return securityLogger }

// Errors returns the error-only logger (file gets LevelError records only).
func Errors() *slog.Logger { return errorLogger }

// Init wires the four log streams. dir is the log directory (default "./logs");
// env is the runtime environment ("dev" → human text + debug, else JSON). Level
// can be overridden with LOG_LEVEL (debug|info|warn|error).
func Init(dir, env string) error {
	if dir == "" {
		dir = "./logs"
	}
	useJSON := env != "dev"
	level := parseLevel(os.Getenv("LOG_LEVEL"), env)

	app, err := stream(dir, "app.log", streamApp, level, useJSON)
	if err != nil {
		return err
	}
	access, err := stream(dir, "access.log", streamAccess, level, useJSON)
	if err != nil {
		return err
	}
	security, err := stream(dir, "security.log", streamSecurity, level, useJSON)
	if err != nil {
		return err
	}
	// Error stream is pinned to LevelError regardless of LOG_LEVEL.
	errs, err := stream(dir, "error.log", streamError, slog.LevelError, useJSON)
	if err != nil {
		return err
	}

	appLogger, accessLogger, securityLogger, errorLogger = app, access, security, errs
	slog.SetDefault(app)
	return nil
}

func stream(dir, file, name string, level slog.Leveler, useJSON bool) (*slog.Logger, error) {
	rw, err := NewRotatingWriter(dir, file, maxLines, keepFile)
	if err != nil {
		return nil, err
	}
	writers = append(writers, rw)
	w := io.MultiWriter(rw, os.Stdout)
	opts := &slog.HandlerOptions{Level: level}
	var h slog.Handler
	if useJSON {
		h = slog.NewJSONHandler(w, opts)
	} else {
		h = slog.NewTextHandler(w, opts)
	}
	return slog.New(h).With("logger", name), nil
}

// Close flushes and closes all stream files. Call on shutdown.
func Close() {
	for _, w := range writers {
		_ = w.Close()
	}
}

func parseLevel(s, env string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	if env == "dev" {
		return slog.LevelDebug
	}
	return slog.LevelInfo
}
