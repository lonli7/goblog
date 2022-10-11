package bootstrap

import (
	"github.com/gorilla/mux"
	"github.com/lonli7/goblog/pkg/route"
	"github.com/lonli7/goblog/routes"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)
	return router
}
