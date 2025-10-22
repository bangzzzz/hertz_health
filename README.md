# hertz_health

A health check middleware for [CloudWeGo Hertz](https://github.com/cloudwego/hertz) framework that provides liveness and readiness probes for your applications.

## Features

- 🚀 Easy to integrate with Hertz applications
- 🔧 Configurable liveness and readiness endpoints
- 🎯 Custom health check functions
- 📊 Standard HTTP status codes (200 for healthy, 503 for unhealthy)
- 🧪 Comprehensive test coverage

## Installation

```bash
go get github.com/bangzzzz/hertz_health
```

## Quick Start

### Basic Usage

```go
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

    // Your application routes
    h.GET("/", func(ctx context.Context, c *app.RequestContext) {
        c.JSON(200, utils.H{"message": "Hello, World!"})
    })

    log.Println("Health endpoints available:")
    log.Println("  Liveness:  http://localhost:8080/health/liveness")
    log.Println("  Readiness: http://localhost:8080/health/readiness")
    
    h.Spin()
}
```

### Custom Configuration

```go
package main

import (
    "context"
    "database/sql"
    "log"

    "github.com/bangzzzz/hertz_health"
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
)

var db *sql.DB // Your database connection

func main() {
    h := server.Default()

    // Custom health check configuration
    config := hertz_health.Config{
        LivenessProbe: func(ctx *app.RequestContext) bool {
            // Check if the application is running properly
            return true
        },
        LivenessEndpoint: "/health/live",
        ReadinessProbe: func(ctx *app.RequestContext) bool {
            // Check database connection and external dependencies
            return checkDatabaseConnection() && checkExternalServices()
        },
        ReadinessEndpoint: "/health/ready",
    }
    
    h.Use(hertz_health.NewWithConfig(config))

    h.Spin()
}

func checkDatabaseConnection() bool {
    if db == nil {
        return false
    }
    return db.Ping() == nil
}

func checkExternalServices() bool {
    // Check external API availability
    // Return true if all external services are available
    return true
}
```

## Configuration

The `Config` struct allows you to customize the health check behavior:

```go
type Config struct {
    LivenessProbe     HealthCheckerFunc // Function to check liveness
    LivenessEndpoint  string            // Endpoint path for liveness probe
    ReadinessProbe    HealthCheckerFunc // Function to check readiness  
    ReadinessEndpoint string            // Endpoint path for readiness probe
}

type HealthCheckerFunc func(*app.RequestContext) bool
```

### Default Configuration

```go
var DefaultConfig = Config{
    LivenessProbe:     func(ctx *app.RequestContext) bool { return true },
    LivenessEndpoint:  "/health/liveness",
    ReadinessProbe:    func(ctx *app.RequestContext) bool { return true },
    ReadinessEndpoint: "/health/readiness",
}
```

## Health Check Types

### Liveness Probe
- **Purpose**: Indicates whether the application is running
- **Default Endpoint**: `/health/liveness`
- **Use Case**: Kubernetes liveness probe, load balancer health checks
- **Response**: 
  - `200 OK` with `{"status": "ok"}` when healthy
  - `503 Service Unavailable` with `{"status": "unavailable"}` when unhealthy

### Readiness Probe
- **Purpose**: Indicates whether the application is ready to serve traffic
- **Default Endpoint**: `/health/readiness`
- **Use Case**: Kubernetes readiness probe, deployment health checks
- **Response**:
  - `200 OK` with `{"status": "ok"}` when ready
  - `503 Service Unavailable` with `{"status": "unavailable"}` when not ready

## API Reference

### Functions

#### `New(config ...Config) app.HandlerFunc`
Creates a new health check middleware with optional configuration.

#### `NewWithConfig(cfg Config) app.HandlerFunc`
Creates a new health check middleware with the specified configuration.

## Testing

Run the tests:

```bash
go test -v
```

## Examples

Check the [example](./example) directory for a complete working example.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
