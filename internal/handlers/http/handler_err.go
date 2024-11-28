package http_handler

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/giicoo/GiicooAuth/pkg/errTools"

	"net/http"
)

func JSONError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	err := models.ErrorResponse{Err: msg}
	json.NewEncoder(w).Encode(err)
}

// Add err and add_err in JSONError
func JSONHandleError(w http.ResponseWriter, err error, add_err error) {
	var apiErr errTools.APIError
	if errors.As(err, &apiErr) {
		status, msg := apiErr.APIError()
		if add_err != nil {
			JSONError(w, status, fmt.Sprintf("%s: %s", msg, add_err))
		} else {
			JSONError(w, status, msg)
		}

	} else {
		status, msg := errTools.ErrInternalError.APIError()
		JSONError(w, status, msg)
	}
}
