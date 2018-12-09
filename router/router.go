package router

import (
	"net/http"

	"github.com/jongschneider/go-project/router/routes/v1"

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

	RegisterRoutes(router, routes.GetRoutes())
	RegisterSubRoutes(router, v1.GetRoutes())

	return &RouteHandler{router}
}

// RegisterRoute registers a route on a *mux.Router
func RegisterRoute(router *mux.Router, route routes.Route) {
	router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
}

// RegisterRoutes registers routes on a *mux.Router
func RegisterRoutes(router *mux.Router, routes routes.Routes) {
	for _, route := range routes {
		RegisterRoute(router, route)
	}
}

// RegisterSubRoute registers a subRoute on a *mux.Router
func RegisterSubRoute(router *mux.Router, path string, subRoute routes.SubRoute) {
	subRouter := router.PathPrefix(path).Subrouter()
	subRouter.Use(subRoute.Middleware)
	for _, route := range subRoute.Routes {
		RegisterRoute(subRouter, route)
	}
}

// RegisterSubRoutes registers subRoutes on a *mux.Router
func RegisterSubRoutes(router *mux.Router, subRoutes routes.SubRoutes) {
	for path, subRoute := range subRoutes {
		RegisterSubRoute(router, path, subRoute)
	}
}
