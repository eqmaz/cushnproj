package health

// --- gRPC service definitions ---

// TODO - we could implement these in protobuf files, but this is just a quickie

// HealthCheckRequest is the gRPC request.
type HealthCheckRequest struct{}

// HealthCheckResponse is the gRPC response.
type HealthCheckResponse struct {
	Status  string
	Message string
}
