package internal

import (
	"context"
	"fmt"
	"net/http"

	"energy_tk/internal/application/users"
	"energy_tk/internal/config"
	"energy_tk/internal/infrastructure/mariadb"
	"energy_tk/internal/infrastructure/rest"
	usersRepo "energy_tk/internal/repositories/users"
	"energy_tk/pkg/clients/energytksite"
)

func RunApp(ctx context.Context) error {
	cfg := config.MustLoad()

	db, err := mariadb.Connect(ctx, cfg.MariaDbDSN)
	if err != nil {
		return err
	}

	if err := mariadb.Migrate(cfg.MariaDbDSN)(db); err != nil {
		return fmt.Errorf("RunApp - migrate error: %w", err)
	}

	siteClient := energytksite.New()
	_ = siteClient

	usersRepository := usersRepo.New(db)

	usersService := users.New(usersRepository)

	if err := provideRest(usersService); err != nil {
		return err
	}

	return nil
}

func provideRest(usersService *users.Service) error {
	mux := rest.RunServer(usersService)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return server.ListenAndServe()
}
