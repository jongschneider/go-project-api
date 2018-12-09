package routes

import (
	"log"
	"net/http"
)

// Login parses email and password to authenticate a user
func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		// validate email
		pw := r.FormValue("pw")
		// validate pass

		log.Println("email:", email)
		log.Println("pw:", pw)
	}
}
