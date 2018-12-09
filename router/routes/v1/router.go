package v1

import (
	"net/http"

	"github.com/jongschneider/go-project/router/routes"
)

// Middleware is the middleware for v1 routes
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes gets the /v1 subroutes
func GetRoutes() routes.SubRoutes {
	subRoutes := routes.SubRoutes{}
	subRoutes.Add("/v1", routes.SubRoute{
		Routes: routes.Routes{
			{
				Name:        "V1HealthRoute",
				Method:      http.MethodGet,
				Pattern:     "/health",
				HandlerFunc: Health(),
			},
		},
		Middleware: Middleware,
	})

	return subRoutes
}
