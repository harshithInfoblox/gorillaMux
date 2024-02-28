package main

import (
	"context"   // Package for handling context
	"flag"      // Package for parsing command-line flags
	"log"       // Package for logging
	"net/http"  // Package for HTTP server and client implementations
	"os"        // Package provides a platform-independent interface to operating system functionality
	"os/signal" // Package for handling signals such as Ctrl+C
	"time"      // Package for handling time-related operations

	"github.com/gorilla/mux" // Package for HTTP request multiplexer (router)
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter() // Creates a new router using Gorilla Mux

	// Add your routes as needed
	// You can define routes using r.HandleFunc or r.Methods as per your requirements

	srv := &http.Server{ // Creates a new HTTP server instance
		Addr:         "0.0.0.0:8080",   // Specifies the network address to listen on (all available network interfaces on port 8080)
		WriteTimeout: time.Second * 15, // Specifies the maximum duration for writing the entire response (15 seconds)
		ReadTimeout:  time.Second * 15, // Specifies the maximum duration for reading the entire request (15 seconds)
		IdleTimeout:  time.Second * 60, // Specifies the maximum duration to wait for the next request when keep-alive is enabled (60 seconds)
		Handler:      r,                // Specifies the handler to invoke, in this case, the Gorilla Mux router
	}

	go func() { // Start the HTTP server in a separate goroutine
		if err := srv.ListenAndServe(); err != nil { // ListenAndServe starts the HTTP server with the provided handler and blocks until the server is shutdown
			log.Println(err) // Log any error that may occur while starting the server
		}
	}()

	c := make(chan os.Signal, 1)   // Create a channel to receive signals (e.g., Ctrl+C)
	signal.Notify(c, os.Interrupt) // Notify the channel when receiving an interrupt signal (e.g., Ctrl+C)

	<-c // Wait for a signal to be received from the channel

	ctx, cancel := context.WithTimeout(context.Background(), wait) // Create a context with a timeout
	defer cancel()                                                 // Defer the cancellation of the context until the main function exits

	if err := srv.Shutdown(ctx); err != nil { // Shutdown the server gracefully
		log.Fatal(err) // Log and exit if there's an error during the shutdown process
	}
	log.Println("shutting down") // Log a message indicating that the server is shutting down
	os.Exit(0)                   // Exit the program with a status code of 0 (indicating success)
}
