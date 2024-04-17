package config

type App struct {
	MariaDbDSN string `env:"MARIADB_DSN"`
}

func MustLoad() *App {
	return &App{}
}
