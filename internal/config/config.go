package config

import "energy_tk/internal/infrastructure/mariadb"

type App struct {
	MariaDbDSN string `env:"MARIADB_DSN"`
}

func MustLoad() *App {
	return &App{
		MariaDbDSN: mariadb.MakeDSN(
			"127.0.0.1:3306",
			"admin",
			"admin",
			"energy_tk",
		),
	}
}
