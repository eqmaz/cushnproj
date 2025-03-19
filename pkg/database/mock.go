package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// MockDatabase - dummy implementation of Database for tests.
type MockDatabase struct {
	// You can add fields to track calls or preset results if needed.
	ExecFunc     func(query string, args ...interface{}) (sql.Result, error)
	QueryFunc    func(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowFunc func(query string, args ...interface{}) *sqlx.Row
}

func (m *MockDatabase) Connect() error {
	// Simulate a successful connection (or track call)
	return nil
}

func (m *MockDatabase) Close() error {
	return nil
}

func (m *MockDatabase) Ping(ctx context.Context) error {
	return nil // always "healthy"
}

func (m *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(query, args...)
	}
	return nil, nil
}

func (m *MockDatabase) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	if m.QueryFunc != nil {
		return m.QueryFunc(query, args...)
	}
	return nil, nil
}

func (m *MockDatabase) QueryRow(query string, args ...interface{}) *sqlx.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(query, args...)
	}
	return &sqlx.Row{} // return an empty Row
}
