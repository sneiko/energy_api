package internal

import (
	"context"
	"log/slog"
	"net/http"

	"energy_tk/internal/application/invoices"
	"energy_tk/internal/application/users"
	"energy_tk/internal/config"
	"energy_tk/internal/infrastructure/mariadb"
	"energy_tk/internal/infrastructure/rest"
	invoicesRepo "energy_tk/internal/repositories/invoices"
	usersRepo "energy_tk/internal/repositories/users"
	"energy_tk/internal/workers"
	"energy_tk/pkg/clients/energytksite"
)

func RunApp(ctx context.Context) error {
	cfg := config.MustLoad()

	db, err := mariadb.Connect(ctx, cfg.MariaDbDSN)
	if err != nil {
		return err
	}

	//if err := mariadb.Migrate(cfg.MariaDbDSN)(db); err != nil {
	//	return fmt.Errorf("RunApp - migrate error: %w", err)
	//}

	siteClient := energytksite.New()

	usersRepository := usersRepo.New(db)
	invoicesRepository := invoicesRepo.New(db)

	usersService := users.New(usersRepository)
	invoicesService := invoices.New(usersRepository, invoicesRepository, siteClient)

	daemons := workers.New(ctx)
	daemons.Add(workers.NewInvoicesCheck(invoicesService))
	daemons.Start()

	if err := provideRest(
		usersService,
		invoicesService,
	); err != nil {
		return err
	}

	return nil
}

func provideRest(usersService *users.Service, invoicesService *invoices.Service) error {
	mux := rest.RunServer(usersService, invoicesService)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	slog.Info("Starting server on :8080")
	return server.ListenAndServe()
}
