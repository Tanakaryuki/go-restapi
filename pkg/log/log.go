package log

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s ,Body: %v", r.Method, r.URL.Path, r.Body)
		next.ServeHTTP(w, r)
	})
}
