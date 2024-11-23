package http_handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
	errTools "github.com/giicoo/GiicooAuth/pkg/err_tools"
	"github.com/sirupsen/logrus"
)

// @Summary      	Generate JWT Token
// @Description  	generate JWT token
// @Tags         	auth
// @Accept			json
// @Produce			json
// @Param			user	body	models.UserRequest	true	"Write Email and Password"
// @Success		 	200		{object}	models.JwtResponse
// @Failure      	400  	{object}  	models.ErrorResponse
// @Failure      	500  	{object}  	models.ErrorResponse
// @Router       	/generate-token/ [post]
func (h *Handler) GenerateJWT(w http.ResponseWriter, r *http.Request) {
	logrus.Info(r.URL)
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

	jwtResponse, err := h.services.GenerateJWT(user.Email, user.Password)
	if err != nil {
		h.log.Error(fmt.Errorf("error with generate jwt service: %s", err))
		JSONHandleError(w, err, nil)
		return
	}

	if err := SendResponse(w, jwtResponse); err != nil {
		h.log.Error(fmt.Errorf("err with send response: %s", err))
		JSONHandleError(w, err, nil)
		return
	}
}

// @Summary      	Check JWT Token
// @Description  	Check JWT token
// @Tags         	auth
// @Accept			json
// @Produce			json
// @Param			user	body	models.JwtRequest	true	"Write JWT token"
// @Success		 	200		{object}	models.UserResponse
// @Failure      	400  	{object}  	models.ErrorResponse
// @Failure      	500  	{object}  	models.ErrorResponse
// @Router       	/check-token/ [post]
func (h *Handler) CheckJWT(w http.ResponseWriter, r *http.Request) {
	logrus.Info(r.URL)
	body := r.Body
	defer body.Close()

	jwtToken := models.JwtRequest{}

	if err := json.NewDecoder(body).Decode(&jwtToken); err != nil {
		h.log.Error(fmt.Errorf("error with json decoder: %s", err))
		JSONHandleError(w, errTools.ErrInvalidJSON, err)
		return
	}

	err := ValidateStructure(jwtToken)
	if err != nil {
		h.log.Error(fmt.Errorf("error with validate struct: %s", err))
		JSONHandleError(w, err, err)
		return
	}

	user, err := h.services.CheckJWT(jwtToken.JwtToken)
	if err != nil {
		h.log.Error(fmt.Errorf("error with check jwt service: %s", err))
		JSONHandleError(w, err, nil)
		return
	}

	if err := SendResponse(w, user); err != nil {
		h.log.Error(fmt.Errorf("err with send response: %s", err))
		JSONHandleError(w, err, nil)
		return
	}
}
