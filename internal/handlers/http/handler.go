package http_handler

import (
	"net/http"

	_ "github.com/giicoo/GiicooAuth/docs"
	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/services"
	"github.com/giicoo/GiicooAuth/pkg/data"
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

func SendResponse(w http.ResponseWriter, r interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	if err := data.ToJSON(r, w); err != nil {
		return err
	}
	return nil
}

//TODO: testing
