package http_handler

import (
	"encoding/json"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Router       /singup/ [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logrus.Info(r.URL)
	body := r.Body
	defer body.Close()

	user := models.UserRequest{}

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		logrus.Error(err)
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	err := h.services.CreateUser(user.Email, user.Password)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "Error with db", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Successful Sing Up"))
}
