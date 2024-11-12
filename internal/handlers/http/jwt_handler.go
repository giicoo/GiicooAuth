package http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GenerateJWT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	body := r.Body
	defer body.Close()

	user := models.UserRequest{}

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	jwtToken, err := h.services.GenerateJWT(user.Email, user.Password)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(jwtToken))
}

func (h *Handler) CheckJWT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	body := r.Body
	defer body.Close()

	jwtToken := models.JwtResponse{}

	if err := json.NewDecoder(body).Decode(&jwtToken); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
	id, email, err := h.services.CheckJWT(jwtToken.JwtToken)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}
	user := models.UserResponse{
		UserId: id,
		Email:  email,
	}
	if err := json.NewEncoder(w).Encode(user); err != nil {
		logrus.Error(err)
		http.Error(w, "Error send jwt", http.StatusInternalServerError)
		return
	}

}
