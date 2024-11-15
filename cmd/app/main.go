package main

import (
	_ "github.com/giicoo/GiicooAuth/docs"
	"github.com/giicoo/GiicooAuth/internal/app"
)

// @title           GiicooAuth
// @version         1.0
// @description     Service for auth

// @host      localhost:8080
// @BasePath  /
func main() {
	app.RunApp()
}
