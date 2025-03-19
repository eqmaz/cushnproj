package database

import (
	"database/sql"
	"fmt"
	"time"

	c "cushon/pkg/console"
	"cushon/pkg/logger"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	cfg    DbConfig       // Database configuration
	SqlxDb *sqlx.DB       // The sqlx database handle
	logger *logger.Logger // Assume Logger is an interface or concrete type for App.Logger
}

// NewMySQL creates a new MySQL database instance (not yet connected).
func NewMySQL(cfg DbConfig, logger *logger.Logger) *MySQL {
	return &MySQL{cfg: cfg, logger: logger}
}

func (m *MySQL) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true",
		m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port, m.cfg.Name)
	// Attempt connection (sqlx.Connect does sqlx.Open + db.Ping)
	sqlxDb, err := sqlx.Connect(m.cfg.Driver, dsn)
	if err != nil {
		m.logger.Errorf("Database connection failed: %v", err)
		return err
	}

	// Set connection pool options
	sqlxDb.SetMaxOpenConns(m.cfg.MaxOpenCon)
	sqlxDb.SetMaxIdleConns(m.cfg.MaxIdleCon)
	sqlxDb.SetConnMaxLifetime(m.cfg.ConnMaxLifetime)

	m.SqlxDb = sqlxDb
	//m.logger.Infof("Connected to %s database %q at %s:%d", m.cfg.Driver, m.cfg.Name, m.cfg.Host, m.cfg.Port)
	c.Successf("Connected to %s database %q at %s:%d", m.cfg.Driver, m.cfg.Name, m.cfg.Host, m.cfg.Port)

	//var err error
	for i := 1; i <= 3; i++ {
		m.SqlxDb, err = sqlx.Connect(m.cfg.Driver, dsn)
		if err == nil {
			break
		}
		m.logger.Errorf("DB connection attempt %d failed: %v", i, err)
		time.Sleep(2 * time.Second) // wait before retry
	}
	if err != nil {
		// All attempts failed
		return fmt.Errorf("could not connect to database: %w", err)
	}

	return nil
}

func (m *MySQL) Close() error {
	if m.SqlxDb != nil {
		return m.SqlxDb.Close()
	}
	return nil
}

func (m *MySQL) Ping() error {
	return m.SqlxDb.Ping()
}

func (m *MySQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := m.SqlxDb.Exec(query, args...)
	if err != nil {
		m.logger.Errorf("Exec query failed: %v, query=%q", err, query)
	}
	return result, err
}
