package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
