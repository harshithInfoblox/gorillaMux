package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define a route handler for the "/foo" endpoint
	r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)

	// Use the CORS Method Middleware
	r.Use(mux.CORSMethodMiddleware(r))

	// Start the HTTP server
	http.ListenAndServe(":8080", r)
}

// Handler function for the "/foo" endpoint
func fooHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Access-Control-Allow-Origin header to allow requests from any origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Handle OPTIONS requests
	if r.Method == http.MethodOptions {
		return
	}

	// Respond with "foo" for other request methods
	w.Write([]byte("foo"))
}
