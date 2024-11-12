package main

import (
	_ "github.com/giicoo/GiicooAuth/docs"
	"github.com/giicoo/GiicooAuth/internal/app"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8080
// @BasePath  /
func main() {
	app.RunApp()
}
