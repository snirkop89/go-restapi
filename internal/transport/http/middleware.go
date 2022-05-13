package http

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// JSONMiddleware sets the required parameters to return a JSON response
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs each request to the server
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
		}).Info("handled request")

		next.ServeHTTP(w, r)
	})
}

// TimeoutMiddleware cancels any request after a specified timeout.
// if no response is returned within that time, the request will fail,
// and return an error to the user
func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
