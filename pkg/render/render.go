package render

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SuccessResponse is a response for successful requests.
type SuccessResponse struct {
	Data any `json:"data"`
}

// ErrorResponse is a response for failed requests.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Json renders a JSON response with escaped HTML.
func Json(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if status != http.StatusOK {
		data = ErrorResponse{Error: fmt.Sprintf("%s", data)}
	}
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	if err := encoder.Encode(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func DecodeJSON[T any](body io.ReadCloser, model *T) error {
	if err := json.NewDecoder(body).Decode(&model); err != nil {
		return err
	}
	return nil
}
