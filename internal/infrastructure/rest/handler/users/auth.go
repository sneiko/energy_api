package users

import (
	"context"
	"fmt"
	"net/http"

	"energy_tk/pkg/render"
)

type AuthService interface {
	Auth(ctx context.Context, token string) error
}

type AuthRequest struct {
	Token string `json:"token"`
}

func Auth(service AuthService) http.HandlerFunc {
	decodeRequest := func(r *http.Request) (*AuthRequest, error) {
		var req AuthRequest
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			return nil, err
		}

		if req.Token == "" {
			return nil, fmt.Errorf("token is required")
		}

		return &req, nil
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := decodeRequest(r)
		if err != nil {
			render.Json(w, http.StatusBadRequest, err)
		}

		err = service.Auth(r.Context(), req.Token)
		if err != nil {
			render.Json(w, http.StatusBadRequest, err)
		}

		render.Json(w, http.StatusOK, req)
	}
}
