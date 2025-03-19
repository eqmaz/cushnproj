package health

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer wraps a gRPC server.
type GRPCServer struct {
	srv    *grpc.Server
	health *Health
	lis    net.Listener
}

// NewGRPCServer - creates a new gRPC server.
func NewGRPCServer(h *Health) *GRPCServer {
	s := grpc.NewServer()
	gs := &GRPCServer{srv: s, health: h}
	RegisterHealthServer(s, gs)
	reflection.Register(s)
	return gs
}

// Check implements the gRPC health check.
func (g *GRPCServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	result, err := g.health.HealthCheck(ctx)
	status := "SERVING"
	if err != nil {
		status = "NOT_SERVING"
	}
	return &HealthCheckResponse{Status: status, Message: result.Message}, nil
}

// Start runs the gRPC server.
func (g *GRPCServer) Start(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("gRPC listen error: %v", err)
	}
	g.lis = lis
	log.Printf("gRPC health server listening on port %s", port)
	if err := g.srv.Serve(lis); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}
}

// Stop gracefully stops the gRPC server.
func (g *GRPCServer) Stop() {
	g.srv.GracefulStop()
}

// HealthServer defines the gRPC health interface.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// RegisterHealthServer registers the health service.
func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}

func _Health_Check_Handler(
	srv interface{},
	ctx context.Context,
	decodeFunc func(interface{}) error,
	interceptor grpc.UnaryServerInterceptor,
) (interface{}, error) {
	req := new(HealthCheckRequest)
	if err := decodeFunc(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, req)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, req, info, handler)
}
