package mariadb

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	migrateV4 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const migrationsPath = "file://./internal/infrastructure/mariadb/migrations"

var once sync.Once

func migrate(db *sql.DB, migrationsPath, dbName string) error {
	var err error

	once.Do(func() {
		if migrationsPath == "" {
			slog.Error("Migrate: migrations path is empty")
		}

		var migrateErr error
		defer func() {
			err = migrateErr
		}()

		driver, migrateErr := mysql.WithInstance(db, &mysql.Config{})
		if migrateErr != nil {
			return
		}

		m, migrateErr := migrateV4.NewWithDatabaseInstance(migrationsPath, dbName, driver)
		if migrateErr != nil {
			return
		}

		migrateErr = m.Up()
		if errors.Is(migrateErr, migrateV4.ErrNoChange) {
			migrateErr = nil
			slog.Info(fmt.Sprintf("migrations %s", migrateV4.ErrNoChange))
			return
		}

		if migrateErr != nil {
			return
		}
	})

	return err
}

func Migrate(dsn string) func(*sql.DB) error {
	parts := strings.Split(dsn, "/")
	return func(db *sql.DB) error {
		return migrate(db, migrationsPath, parts[1])
	}
}
