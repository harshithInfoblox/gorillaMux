package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"strings"

	"github.com/gorilla/mux"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "./src/mywebpage/static", "the directory to serve files from. Defaults to the 'static' directory")
	flag.Parse()

	r := mux.NewRouter()

	// Custom file server handler to filter out unwanted files and directories
	fs := http.FileServer(http.Dir(dir))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "..") || strings.HasPrefix(r.URL.Path, "/.") {
			http.NotFound(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	})))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
