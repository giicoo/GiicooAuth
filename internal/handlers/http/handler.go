package http_handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/giicoo/GiicooAuth/docs"
	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/giicoo/GiicooAuth/internal/services"
	errTools "github.com/giicoo/GiicooAuth/pkg/err_tools"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	cfg      *config.Config
	log      *logrus.Logger
	services *services.Services
}

func NewHandler(cfg *config.Config, log *logrus.Logger, services *services.Services) *Handler {
	return &Handler{
		cfg:      cfg,
		log:      log,
		services: services,
	}
}

func (h *Handler) CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/singup", h.CreateUser).Methods("POST")
	r.HandleFunc("/generate-token", h.GenerateJWT).Methods("POST")
	r.HandleFunc("/check-token", h.CheckJWT).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	return r
}

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

func ValidateStructure(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return errTools.WrapError(err, errTools.ErrInvalidJSON)
	}
	return nil
}

func SendResponse(w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		return err
	}
	return nil
}

//TODO:handle error
//TODO: testing
//TODO: swagger
