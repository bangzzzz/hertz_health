package main

import (
	"context"
	"log"

	"github.com/bangzzzz/hertz_health"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func main() {
	h := server.Default()

	// Use default health check configuration
	h.Use(hertz_health.New())

	// Or use custom configuration
	customConfig := hertz_health.Config{
		LivenessProbe: func(ctx *app.RequestContext) bool {
			// Custom liveness check logic
			// For example, check if the application is running properly
			return true
		},
		LivenessEndpoint: "/health/live",
		ReadinessProbe: func(ctx *app.RequestContext) bool {
			// Custom readiness check logic
			// For example, check database connection, external services, etc.
			return checkDatabaseConnection() && checkExternalServices()
		},
		ReadinessEndpoint: "/health/ready",
	}
	
	// Apply custom health check middleware
	h.Use(hertz_health.NewWithConfig(customConfig))

	// Regular application routes
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{
			"message": "Hello, World!",
		})
	})

	h.GET("/api/users", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(200, utils.H{
			"users": []string{"Alice", "Bob", "Charlie"},
		})
	})

	log.Println("Server starting on :8080")
	log.Println("Health endpoints:")
	log.Println("  Liveness:  http://localhost:8080/health/live")
	log.Println("  Readiness: http://localhost:8080/health/ready")
	
	h.Spin()
}

// Mock functions for demonstration
func checkDatabaseConnection() bool {
	// Simulate database connection check
	// In real application, you would check actual database connection
	return true
}

func checkExternalServices() bool {
	// Simulate external services check
	// In real application, you would check external API availability
	return true
}