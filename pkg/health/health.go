package health

import (
	"context"
	"sync/atomic"
)

// HealthResult represents the health check outcome.
type HealthResult struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// HealthCheckFunc defines the business logic for health.
type HealthCheckFunc func(ctx context.Context) (HealthResult, error)

// Options configures the health server.
type Options struct {
	HTTPEnabled   bool
	GRPCEnabled   bool
	HTTPPort      string
	GRPCPort      string
	CheckFunction HealthCheckFunc
}

// Health is the singleton health object.
type Health struct {
	opts      Options
	callCount uint64
	httpSrv   *HTTPServer
	grpcSrv   *GRPCServer
}

// New creates a new Health instance.
func New(opts Options) *Health {
	return &Health{opts: opts}
}

// HealthCheck invokes the health check logic and increments the counter.
func (h *Health) HealthCheck(ctx context.Context) (HealthResult, error) {
	atomic.AddUint64(&h.callCount, 1)
	return h.opts.CheckFunction(ctx)
}

// GetCallCount returns the total number of health check calls.
func (h *Health) GetCallCount() uint64 {
	return atomic.LoadUint64(&h.callCount)
}

// Start spins up the enabled servers.
func (h *Health) Start() {
	if h.opts.HTTPEnabled {
		h.httpSrv = NewHTTPServer(h)
		go h.httpSrv.Start(h.opts.HTTPPort)
	}
	if h.opts.GRPCEnabled {
		h.grpcSrv = NewGRPCServer(h)
		go h.grpcSrv.Start(h.opts.GRPCPort)
	}
}

// Stop gracefully stops any running servers.
func (h *Health) Stop() {
	if h.opts.HTTPEnabled && h.httpSrv != nil {
		h.httpSrv.Stop()
	}
	if h.opts.GRPCEnabled && h.grpcSrv != nil {
		h.grpcSrv.Stop()
	}
}
