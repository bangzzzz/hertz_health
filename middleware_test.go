package hertz_health

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestHealthCheckMiddleware_DefaultConfig(t *testing.T) {
	h := server.Default()
	h.Use(New())

	// Test liveness endpoint
	w := ut.PerformRequest(h.Engine, "GET", "/health/liveness", nil)
	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "ok"))
}

func TestHealthCheckMiddleware_CustomConfig(t *testing.T) {
	config := Config{
		LivenessProbe:     func(ctx *app.RequestContext) bool { return true },
		LivenessEndpoint:  "/custom/live",
		ReadinessProbe:    func(ctx *app.RequestContext) bool { return false },
		ReadinessEndpoint: "/custom/ready",
	}

	h := server.Default()
	h.Use(NewWithConfig(config))

	// Test custom liveness endpoint (should return OK)
	w1 := ut.PerformRequest(h.Engine, "GET", "/custom/live", nil)
	assert.DeepEqual(t, http.StatusOK, w1.Code)
	assert.True(t, strings.Contains(w1.Body.String(), "ok"))

	// Test custom readiness endpoint (should return Service Unavailable)
	w2 := ut.PerformRequest(h.Engine, "GET", "/custom/ready", nil)
	assert.DeepEqual(t, http.StatusServiceUnavailable, w2.Code)
	assert.True(t, strings.Contains(w2.Body.String(), "unavailable"))
}

func TestHealthCheckMiddleware_NonHealthEndpoint(t *testing.T) {
	h := server.Default()
	h.Use(New())
	
	// Add a regular route
	h.GET("/api/test", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, map[string]string{"message": "test"})
	})

	w := ut.PerformRequest(h.Engine, "GET", "/api/test", nil)
	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "test"))
}

func TestHealthCheckMiddleware_NilChecker(t *testing.T) {
	config := Config{
		LivenessProbe:     nil,
		LivenessEndpoint:  "/health/liveness",
		ReadinessProbe:    nil,
		ReadinessEndpoint: "/health/readiness",
	}

	h := server.Default()
	h.Use(NewWithConfig(config))

	// Test with nil checker (should default to OK)
	w := ut.PerformRequest(h.Engine, "GET", "/health/liveness", nil)
	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "ok"))
}