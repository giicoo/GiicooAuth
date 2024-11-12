package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/giicoo/GiicooAuth/internal/config"
	http_handler "github.com/giicoo/GiicooAuth/internal/handlers/http"
	"github.com/giicoo/GiicooAuth/internal/repository/sqlite"
	"github.com/giicoo/GiicooAuth/internal/server"
	"github.com/giicoo/GiicooAuth/internal/services"
	"github.com/giicoo/GiicooAuth/pkg/log_tool"
	_ "github.com/mattn/go-sqlite3"
)

func RunApp() {

	log := log_tool.NewLogTool()
	// Load Config
	cfg, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Config load.")

	// Load DB
	db, err := sql.Open("sqlite3", cfg.DB.Path)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("DB load.")

	// Init Layers
	repo := sqlite.NewRepo(cfg, log, db)
	services := services.NewServices(cfg, log, repo)
	handler := http_handler.NewHandler(cfg, log, services)

	router := handler.CreateRouter()
	if err = repo.InitDB(); err != nil {
		log.Fatal(err)
	}

	// Start Server
	srv := server.NewServer(cfg, router)

	go func() {

		err := srv.StartServer()
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				fmt.Println()
			default:
				log.Fatal(err)
			}
		}
	}()
	log.Info("Server start.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	// ShutDown Server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.ShutdownServer(ctx); err != nil {
		log.Error(err)
	} else {
		log.Info("Server stop.")
	}

	err = db.Close()
	if err != nil {
		log.Error(err)
	}
	log.Info("DB close.")
}
