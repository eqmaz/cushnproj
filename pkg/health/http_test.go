// File: http_test.go
package health

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHealthSuccess(t *testing.T) {
	// Dummy check function that returns success.
	checkFunc := func(ctx context.Context) (HealthResult, error) {
		return HealthResult{Status: "OK", Message: "healthy"}, nil
	}
	opts := Options{CheckFunction: checkFunc}
	h := New(opts)
	srv := NewHTTPServer(h)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := srv.app.Test(req)
	if err != nil {
		t.Fatalf("HTTP test failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
	var result HealthResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if result.Status != "OK" {
		t.Errorf("expected OK status, got %s", result.Status)
	}
}

func TestHTTPHealthFailure(t *testing.T) {
	// Dummy check function that returns failure.
	checkFunc := func(ctx context.Context) (HealthResult, error) {
		return HealthResult{Status: "FAIL", Message: "unhealthy"}, errors.New("failure")
	}
	opts := Options{CheckFunction: checkFunc}
	h := New(opts)
	srv := NewHTTPServer(h)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := srv.app.Test(req)
	if err != nil {
		t.Fatalf("HTTP test failed: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected status 503, got %d", resp.StatusCode)
	}
	var result HealthResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if result.Status != "FAIL" {
		t.Errorf("expected FAIL status, got %s", result.Status)
	}
}
