// Package health exposes a generic /health handler wired around a set of
// user-defined checkers (database, cache, downstream services, ...). Each
// service registers whatever dependencies matter for its liveness contract.
package health

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Checker is the contract each dependency implements. Name is stable and
// surfaced in the JSON payload; Check runs against the supplied ctx.
type Checker interface {
	Name() string
	Check(ctx context.Context) error
}

// CheckerFunc adapts a plain function into a Checker.
type CheckerFunc struct {
	N string
	F func(ctx context.Context) error
}

func (c CheckerFunc) Name() string                        { return c.N }
func (c CheckerFunc) Check(ctx context.Context) error     { return c.F(ctx) }

// Config configures a Handler. Location/Timezone are informational only; they
// echo in the response so operators can spot region drift.
type Config struct {
	Version  string
	Location string
	Timezone string
	Timeout  time.Duration // per-checker timeout; defaults to 2s
	Checkers []Checker
}

type Handler struct {
	cfg       Config
	startTime time.Time
}

func New(cfg Config) *Handler {
	if cfg.Timeout == 0 {
		cfg.Timeout = 2 * time.Second
	}
	return &Handler{cfg: cfg, startTime: time.Now()}
}

// Gin serves GET /health. Returns 200 when every checker succeeds, otherwise 503.
func (h *Handler) Gin(c *gin.Context) {
	checks := gin.H{}
	allUp := true
	for _, ck := range h.cfg.Checkers {
		ctx, cancel := context.WithTimeout(c.Request.Context(), h.cfg.Timeout)
		err := ck.Check(ctx)
		cancel()
		up := err == nil
		if !up {
			allUp = false
		}
		checks[ck.Name()] = gin.H{"status": statusFor(up)}
	}

	status := "UP"
	code := http.StatusOK
	if !allUp {
		status = "DOWN"
		code = http.StatusServiceUnavailable
	}

	c.JSON(code, gin.H{
		"status": status,
		"environment": gin.H{
			"location":     h.cfg.Location,
			"timezone":     h.cfg.Timezone,
			"current_time": time.Now().Format(time.RFC3339),
			"uptime":       time.Since(h.startTime).String(),
		},
		"version": h.cfg.Version,
		"checks":  checks,
	})
}

func statusFor(up bool) string {
	if up {
		return "UP"
	}
	return "DOWN"
}
