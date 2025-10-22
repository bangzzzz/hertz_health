package hertz_health

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// New creates a new health check middleware with the given configuration
func New(config ...Config) app.HandlerFunc {
	cfg := DefaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	return func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Path())

		// Handle liveness probe
		if path == cfg.LivenessEndpoint {
			handleHealthCheck(ctx, c, cfg.LivenessProbe)
			return
		}

		// Handle readiness probe
		if path == cfg.ReadinessEndpoint {
			handleHealthCheck(ctx, c, cfg.ReadinessProbe)
			return
		}

		// Continue to next handler if not a health check endpoint
		c.Next(ctx)
	}
}

// handleHealthCheck executes the health check function and returns appropriate response
func handleHealthCheck(ctx context.Context, c *app.RequestContext, checker HealthCheckerFunc) {
	if checker == nil {
		c.JSON(http.StatusOK, utils.H{
			"status": "ok",
		})
		return
	}

	if checker(c) {
		c.JSON(http.StatusOK, utils.H{
			"status": "ok",
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, utils.H{
			"status": "unavailable",
		})
	}
}

// NewWithConfig creates a new health check middleware with custom configuration
func NewWithConfig(cfg Config) app.HandlerFunc {
	return New(cfg)
}