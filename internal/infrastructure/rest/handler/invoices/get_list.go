package invoices

import (
	"context"
	"net/http"

	"energy_tk/internal/domain"
	"energy_tk/internal/infrastructure/rest/middleware"
	"energy_tk/pkg/render"
)

type GetInvoiceListService interface {
	GetListByUserToken(ctx context.Context, token string) (domain.InvoiceList, error)
}

func GetInvoiceList(service GetInvoiceListService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := middleware.GetUserToken(r.Context())

		list, err := service.GetListByUserToken(r.Context(), token)
		if err != nil {
			render.Json(w, http.StatusBadRequest, err)
		}

		render.Json(w, http.StatusOK, list)
	}
}
