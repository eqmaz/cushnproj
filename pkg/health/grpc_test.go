package health

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
)

// HealthClient is a minimal client interface to test our gRPC service.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/health.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// startTestGRPCServer starts a gRPC server on a random port and returns its address and a cleanup function.
func startTestGRPCServer(h *Health, port string) (*GRPCServer, string, func()) {
	srv := NewGRPCServer(h)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv.lis = lis
	go func() {
		if err := srv.srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return srv, lis.Addr().String(), func() {
		srv.Stop()
		err := lis.Close()
		if err != nil {
			// Should not happen, but log just in case.
			log.Printf("failed to close listener: %v", err)
			return
		}
	}
}

func TestGRPCHealthSuccess(t *testing.T) {
	// Dummy check function that returns success.
	checkFunc := func(ctx context.Context) (HealthResult, error) {
		return HealthResult{Status: "OK", Message: "healthy"}, nil
	}
	opts := Options{CheckFunction: checkFunc}
	h := New(opts)
	_, addr, cleanup := startTestGRPCServer(h, "0")
	defer cleanup()

	// Give the server a moment to start.
	time.Sleep(50 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			// This should not happen, but log just in case.
			println(err.Error())
			return
		}
	}(conn)

	client := NewHealthClient(conn)
	resp, err := client.Check(context.Background(), &HealthCheckRequest{})
	if err != nil {
		t.Fatalf("gRPC Check failed: %v", err)
	}
	if resp.Status != "SERVING" {
		t.Errorf("expected SERVING, got %s", resp.Status)
	}
}

func TestGRPCHealthFailure(t *testing.T) {
	// Dummy check function that returns failure.
	checkFunc := func(ctx context.Context) (HealthResult, error) {
		return HealthResult{Status: "FAIL", Message: "unhealthy"}, errors.New("failure")
	}
	opts := Options{CheckFunction: checkFunc}
	h := New(opts)
	_, addr, cleanup := startTestGRPCServer(h, "0")
	defer cleanup()

	time.Sleep(50 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			// This should not happen, but log just in case.
			println(err.Error())
		}
	}(conn)

	client := NewHealthClient(conn)
	resp, err := client.Check(context.Background(), &HealthCheckRequest{})
	if err != nil {
		t.Fatalf("gRPC Check failed: %v", err)
	}
	if resp.Status != "NOT_SERVING" {
		t.Errorf("expected NOT_SERVING, got %s", resp.Status)
	}
}
