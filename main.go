package main

import (
	"github.com/gorilla/mux"
	"github.com/m21power/GrinGram/controllers/routes"
	_ "github.com/m21power/GrinGram/docs"
)

// @title GrinGram API
// @version 1.0
// @description API documentation for GrinGram.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	route := mux.NewRouter()
	r := routes.NewRouter(route)
	r.RegisterRoute()
	r.Run(":8080", route)
}
