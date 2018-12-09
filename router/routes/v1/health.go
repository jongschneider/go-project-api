package v1

import "net/http"

// Health returns a healthCheck handlerFunc
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("v1 is good."))
	}
}
