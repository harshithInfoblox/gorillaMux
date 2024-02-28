package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/products/").Subrouter()
	s.HandleFunc("/{key}", ProductHandler)
	s.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	//r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", s))
}
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "product: %v \n", vars["key"])

}
