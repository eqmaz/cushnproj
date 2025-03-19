package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Database is the abstraction for a database connection.
type Database interface {
	Connect() error                 // Initialize the connection (with pooling)
	Close() error                   // Close the connection, releasing resources
	Ping(ctx context.Context) error // Check if the DB is reachable

	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(query string, args ...interface{}) *sqlx.Row
}
