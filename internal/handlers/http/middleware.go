package http_handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/giicoo/GiicooAuth/pkg/data"
	"github.com/giicoo/GiicooAuth/pkg/errTools"
)

func (h *Handler) MiddlewareValidateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()

		user := models.UserRequest{}

		if err := data.FromJSON(&user, body); err != nil {
			h.log.Error(fmt.Errorf("error with json decoder: %s", err))
			JSONHandleError(w, errTools.ErrInvalidJSON, err)
			return
		}

		err := data.ValidateStructure(user)
		if err != nil {
			h.log.Error(fmt.Errorf("error with validate struct: %s", err))
			JSONHandleError(w, err, err)
			return
		}

		// add the user to the context
		ctx := context.WithValue(r.Context(), userKey{}, user)
		r = r.WithContext(ctx)

		// call the next handler
		next.ServeHTTP(w, r)
	})
}
