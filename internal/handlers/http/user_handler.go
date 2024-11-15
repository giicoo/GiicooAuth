package http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
	errTools "github.com/giicoo/GiicooAuth/pkg/err_tools"
)

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

	user := models.UserRequest{}

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		h.log.Error(fmt.Errorf("error with json decoder: %s", err))
		JSONHandleError(w, errTools.ErrInvalidJSON, err)
		return
	}

	err := ValidateStructure(user)
	if err != nil {
		h.log.Error(fmt.Errorf("error with validate struct: %s", err))
		JSONHandleError(w, err, err)
		return
	}

	userResponse, err := h.services.CreateUser(user.Email, user.Password)
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
