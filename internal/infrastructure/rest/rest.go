package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"energy_tk/internal/application/users"
	usersHandler "energy_tk/internal/infrastructure/rest/handler/users"
)

func RunServer(usersService *users.Service) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/auth", usersHandler.Auth(usersService))
		})
		r.Route("/invoices", func(r chi.Router) {
			r.Get("/list", nil)
			r.Get("/{id}", nil)
		})
	})

	return router
}
