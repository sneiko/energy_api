package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	//nolint:revive,nolintlint // Is needed for correct driver work.
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
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
func Connect(ctx context.Context, dsn string, options ...bun.DBOption) (*bun.DB, error) {
	db, dbErr := sql.Open(mysqlDriver, dsn)
	if dbErr != nil {
		return nil, fmt.Errorf("db connect: %w", dbErr)
	}

	dbConn := bun.NewDB(db, mysqldialect.New(), options...)

	dbConn.SetMaxOpenConns(maxOpenConns)
	dbConn.SetMaxIdleConns(maxIdleConns)
	dbConn.SetConnMaxIdleTime(connMaxIdleTime)
	dbConn.SetConnMaxLifetime(connMaxLifetime)

	if err := dbConn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("db connect: %w", err)
	}

	return dbConn, nil
}

// MakeDSN make data source name.
func MakeDSN(addr, username, password, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, addr, database)
}
