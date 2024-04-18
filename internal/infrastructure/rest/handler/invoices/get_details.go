package invoices

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"energy_tk/internal/domain"
	"energy_tk/pkg/render"
)

type GetInvoiceDetailService interface {
	GetDetails(ctx context.Context, id int) (*domain.Invoice, error)
}

func GetInvoiceDetail(service GetInvoiceDetailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			render.Json(w, http.StatusBadRequest, "id is required")
			return
		}

		list, err := service.GetDetails(r.Context(), id)
		if err != nil {
			render.Json(w, http.StatusBadRequest, err)
		}

		render.Json(w, http.StatusOK, list)
	}
}
