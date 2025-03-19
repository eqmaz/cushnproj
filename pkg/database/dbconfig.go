package database

import "time"

// DbConfig holds database connection parameters.
type DbConfig struct {
	Driver   string // "mysql" | "postgres"
	Host     string
	Port     int
	User     string
	Password string
	Name     string // Schema name

	// Connection pool settings:
	MaxOpenCon      int
	MaxIdleCon      int
	ConnMaxLifetime time.Duration
}
