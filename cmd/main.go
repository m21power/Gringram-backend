package main

import (
	"github.com/gorilla/mux"
	"github.com/m21power/GrinGram/controllers/routes"
)

func main() {
	route := mux.NewRouter()
	r := routes.NewRouter(route)
	r.RegisterRoute()
	r.Run(":8080", route)
}
