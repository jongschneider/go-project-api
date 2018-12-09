package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jongschneider/go-project/router/routes"
)

const staticDir = "/static/"

// RouteHandler represents a mux router
type RouteHandler struct {
	Router *mux.Router
}

// New creates a new router
func New() *RouteHandler {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	router.Use(routes.Middleware)

	for _, route := range routes.GetRoutes() {
		RegisterRoute(router, route)
	}

	return &RouteHandler{
		Router: router,
	}
}

// RegisterRoute registers a route on a *mux.Router
func RegisterRoute(router *mux.Router, route routes.Route) {
	router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.Handler)
}
