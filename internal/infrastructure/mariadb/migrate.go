package mariadb

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	migrateV4 "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

const migrationsPath = "file://./internal/infrastructure/mariadb/migrations"

var once sync.Once

func migrate(db *bun.DB, migrationsPath, dbName string) error {
	var err error

	once.Do(func() {
		if migrationsPath == "" {
			slog.Error("Migrate: migrations path is empty")
		}

		var migrateErr error
		defer func() {
			err = migrateErr
		}()

		db := bun.NewDB(db.DB, mysqldialect.New())

		driver, migrateErr := mysql.WithInstance(db.DB, &mysql.Config{})
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

func Migrate(dsn string) func(*bun.DB) error {
	parts := strings.Split(dsn, "/")
	return func(db *bun.DB) error {
		return migrate(db, migrationsPath, parts[1])
	}
}
