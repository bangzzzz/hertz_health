package hertz_health

import "github.com/cloudwego/hertz/pkg/app"

type HealthCheckerFunc func(*app.RequestContext) bool

type Config struct {
	LivenessProbe HealthCheckerFunc
	LivenessEndpoint string

	ReadinessProbe HealthCheckerFunc
	ReadinessEndpoint string
}

var DefaultConfig = Config{
	LivenessProbe:    func(ctx *app.RequestContext) bool { return true },
	LivenessEndpoint: "/health/liveness",
	ReadinessProbe:   func(ctx *app.RequestContext) bool { return true },
	ReadinessEndpoint: "/health/readiness",
}

