// File: health_test.go
package health

import (
	"context"
	"errors"
	"testing"
)

func TestHealthCallCount(t *testing.T) {
	var calls int
	checkFunc := func(ctx context.Context) (HealthResult, error) {
		calls++
		return HealthResult{Status: "OK", Message: "all good"}, nil
	}
	opts := Options{
		CheckFunction: checkFunc,
	}
	h := New(opts)

	// Invoke the health check a few times
	for i := 0; i < 3; i++ {
		if _, err := h.HealthCheck(context.Background()); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	}
	if h.GetCallCount() != 3 {
		t.Errorf("expected call count 3, got %d", h.GetCallCount())
	}
	if calls != 3 {
		t.Errorf("expected underlying function to be called 3 times, got %d", calls)
	}
}

func TestHealthCheckError(t *testing.T) {
	checkFunc := func(ctx context.Context) (HealthResult, error) {
		return HealthResult{Status: "FAIL", Message: "error occurred"}, errors.New("failure")
	}
	opts := Options{
		CheckFunction: checkFunc,
	}
	h := New(opts)
	result, err := h.HealthCheck(context.Background())
	if err == nil {
		t.Fatal("expected error but got nil")
	}
	if result.Status != "FAIL" {
		t.Errorf("expected FAIL status, got %s", result.Status)
	}
}
