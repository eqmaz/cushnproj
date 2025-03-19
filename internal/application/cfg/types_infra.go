package cfg

// Infrastructure configuration types
// These define behaviour of the application framework

// Database - configuration for the relational database connection
type Database struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"` // The schema name
	User string `json:"db_user"`
	Pass string `json:"db_pass"`
}

// Console - behaviour for CLI console messages
// In production these would be off, and only the logger would be used
type Console struct {
	Enabled bool `json:"enabled"`
	Color   bool `json:"colors"`
}

// Logger - default behaviour for the app-level logger
// This does not include contextual loggers for REST and gRPC requests.
type Logger struct {
	Enabled bool   `json:"enabled"`
	Level   string `json:"level"`
	Stream  string `json:"stream"`
}

// RateLimit - rate limiting configuration for HTTP requests
// We would expect the gateway to handle this, but it's good to have app-level protection as well
type RateLimit struct {
	Enabled     bool `json:"enabled"`
	MaxRequests int  `json:"max_requests"`
	Timeframe   int  `json:"timeframe"`
}

// Health - configuration for the health check endpoints
// We can enable HTTP and/or gRPC based heartbeat listeners
type Health struct {
	Http struct {
		Enabled bool `json:"enabled"`
		Port    int  `json:"port"`
	} `json:"http"`
	Grpc struct {
		Enabled bool `json:"enabled"`
		Port    int  `json:"port"`
	} `json:"grpc"`
}
