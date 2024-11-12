package http_handler

import (
	_ "github.com/giicoo/GiicooAuth/docs"
	"github.com/giicoo/GiicooAuth/internal/config"
	"github.com/giicoo/GiicooAuth/internal/services"
	"github.com/julienschmidt/httprouter"
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

func (h *Handler) CreateRouter() *httprouter.Router {
	router := httprouter.New()
	router.POST("/singup", h.CreateUser)
	router.POST("/generate-token", h.GenerateJWT)
	router.POST("/check-token", h.CheckJWT)

	router.HandlerFunc("GET", "/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/swagger.json")))

	return router
}

//TODO:handle error
//TODO: testing
//TODO: swagger
