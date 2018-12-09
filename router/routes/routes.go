package routes

import (
	"net/http"
)

// Routes is a collection of Route
type Routes []Route

// Route represents a route for the api
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// SubRoute is a collection of Routes and middleware for a subRoute
// 	ex: "/v1" subroute
type SubRoute struct {
	Routes
	Middleware func(http.Handler) http.Handler
}

// SubRoutes is a collection of SubRoutes
type SubRoutes map[string]SubRoute

// Add adds a SubRoute to SubRoutes
func (s SubRoutes) Add(path string, subRoute SubRoute) {
	s[path] = subRoute
}

// Middleware is the main middleware for the application
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes get the routes
func GetRoutes() Routes {
	return Routes{
		Route{
			Name:        "HealthCheck",
			Method:      http.MethodGet,
			Pattern:     "/health",
			HandlerFunc: Health(),
		},
	}
}
