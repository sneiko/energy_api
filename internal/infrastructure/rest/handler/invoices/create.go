package invoices

import (
	"context"
	"fmt"
	"net/http"

	"energy_tk/internal/infrastructure/rest/middleware"
	"energy_tk/pkg/render"
)

type CreateInvoiceService interface {
	Create(ctx context.Context, token string, number string) error
}

type CreateInvoiceRequest struct {
	Number string `json:"number"`
}

func CreateInvoice(service CreateInvoiceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := middleware.GetUserToken(r.Context())
		req, err := render.DecodeJSON[CreateInvoiceRequest](r.Body)
		if err != nil {
			render.Json(w, http.StatusBadRequest, fmt.Errorf("invalid request: %w", err))
		}

		if err := service.Create(r.Context(), token, req.Number); err != nil {
			render.Json(w, http.StatusBadRequest, err)
		}

		render.Json(w, http.StatusOK, nil)
	}
}
