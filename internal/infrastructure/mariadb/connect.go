package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	//nolint:revive,nolintlint // Is needed for correct driver work.
	_ "github.com/go-sql-driver/mysql"
)

const (
	// mysqlDriver name of pg driver.
	mysqlDriver = "mysql"
)

const (
	maxOpenConns    = 3
	maxIdleConns    = 1
	connMaxIdleTime = time.Minute * 10
	connMaxLifetime = time.Minute * 5
)

// Options set any sql.DB properties.
type Options func(*sql.DB) *sql.DB

// Connect return new db connection pool.
func Connect(ctx context.Context, dsn string) (*sql.DB, error) {
	db, dbErr := sql.Open(mysqlDriver, dsn)
	if dbErr != nil {
		return nil, fmt.Errorf("db connect: %w", dbErr)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime)
	db.SetConnMaxLifetime(connMaxLifetime)

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}

	return db, nil
}

// MakeDSN make data source name.
func MakeDSN(addr, username, password, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, addr, database)
}
