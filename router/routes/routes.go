package routes

import (
	"net/http"
)

// Routes is a collection of Route
type Routes []Route

// Route represents a route for the api
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type SubRoute struct {
	Routes
	Middleware func(http.Handler) http.Handler
}

// Middleware is the main middleware for the application
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() Routes {
	return Routes{
		Route{
			Name:    "HealthCheck",
			Method:  http.MethodGet,
			Pattern: "/health",
			Handler: http.HandlerFunc(Health()),
		},
	}
}
