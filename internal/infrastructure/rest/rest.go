package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	invoicesService "energy_tk/internal/application/invoices"
	usersService "energy_tk/internal/application/users"
	"energy_tk/internal/infrastructure/rest/handler/invoices"
	usersHandler "energy_tk/internal/infrastructure/rest/handler/users"
	mw "energy_tk/internal/infrastructure/rest/middleware"
)

func RunServer(
	usersService *usersService.Service,
	invoicesService *invoicesService.Service,
) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/auth", usersHandler.Auth(usersService))
		})

		// With Auth
		r.With(mw.AppAuthMiddleware()).
			Group(func(r chi.Router) {
				r.Route("/invoices", func(r chi.Router) {
					r.Post("/add", invoices.CreateInvoice(invoicesService))
					r.Get("/list", invoices.GetInvoiceList(invoicesService))
					r.Get("/{id}", invoices.GetInvoiceDetail(invoicesService))
				})
			})

	})

	return router
}
