package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Handler for the root endpoint
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

// Middleware for logging incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log details of the incoming request
		log.Printf("Incoming request: %s %s\n", r.Method, r.URL.Path)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Register routes with the router
	r.HandleFunc("/", handler)

	// Add the logging middleware to the router
	r.Use(loggingMiddleware)

	// Create a HTTP server with the router as the handler
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start the HTTP server
	log.Println("Server listening on port 8080...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
