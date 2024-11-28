package http_handler

import (
	"fmt"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
)

type userKey struct{}

// @Summary      	Create User
// @Description  	create user
// @Tags         	users
// @Accept			json
// @Produce			json
// @Param			user	body	models.UserRequest	true	"Write Email and Password"
// @Success		 	200		{object}	models.UserResponse
// @Failure      	400  	{object}  	models.ErrorResponse
// @Failure      	500  	{object}  	models.ErrorResponse
// @Router       	/singup/ [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.log.Info(r.URL)
	body := r.Body
	defer body.Close()

	user := r.Context().Value(userKey{}).(models.UserRequest)

	userResponse, err := h.services.UserService.CreateUser(user.Email, user.Password)
	if err != nil {
		h.log.Error(fmt.Errorf("error with service create user: %s", err))
		JSONHandleError(w, err, nil)
		return
	}

	if err := SendResponse(w, userResponse); err != nil {
		h.log.Error(fmt.Errorf("err with send response: %s", err))
		JSONHandleError(w, err, nil)
		return
	}
}
