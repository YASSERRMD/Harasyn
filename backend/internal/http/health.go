package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type HealthStatus struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Checks    map[string]string `json:"checks"`
}

type HealthChecker struct {
	checks map[string]func(ctx context.Context) error
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{
		checks: make(map[string]func(ctx context.Context) error),
	}
}

func (h *HealthChecker) RegisterCheck(name string, check func(ctx context.Context) error) {
	h.checks[name] = check
}

func (h *HealthChecker) CheckHealth(ctx context.Context) *HealthStatus {
	status := &HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now(),
		Checks:    make(map[string]string),
	}

	for name, check := range h.checks {
		if err := check(ctx); err != nil {
			status.Checks[name] = fmt.Sprintf("unhealthy: %v", err)
			status.Status = "unhealthy"
		} else {
			status.Checks[name] = "healthy"
		}
	}

	return status
}

func (h *HealthChecker) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		health := h.CheckHealth(ctx)

		w.Header().Set("Content-Type", "application/json")
		if health.Status == "unhealthy" {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(health)
	})
}

type ReadinessChecker struct {
	dependencies []string
	checkFuncs   map[string]func(ctx context.Context) error
}

func NewReadinessChecker() *ReadinessChecker {
	return &ReadinessChecker{
		checkFuncs: make(map[string]func(ctx context.Context) error),
	}
}

func (r *ReadinessChecker) AddDependency(name string, check func(ctx context.Context) error) {
	r.dependencies = append(r.dependencies, name)
	r.checkFuncs[name] = check
}

func (r *ReadinessChecker) CheckReadiness(ctx context.Context) (bool, map[string]string) {
	results := make(map[string]string)
	allReady := true

	for _, dep := range r.dependencies {
		if err := r.checkFuncs[dep](ctx); err != nil {
			results[dep] = fmt.Sprintf("not ready: %v", err)
			allReady = false
		} else {
			results[dep] = "ready"
		}
	}

	return allReady, results
}

func (r *ReadinessChecker) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
		defer cancel()

		ready, deps := r.CheckReadiness(ctx)

		w.Header().Set("Content-Type", "application/json")
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":       "ready",
			"dependencies": deps,
		})
	})
}
