package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Article Page: %v\n", mux.Vars(r))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("article")

	// Building URL
	url, err := r.Get("article").URL("category", "technology", "id", "42")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Built URL:", url.String()) // Output: "/articles/technology/42"

	// Additional route for demonstration
	r.HandleFunc("/products/{category}/{id}", ArticleHandler).Name("product")

	// Building URL for the additional route
	productURL, err := r.Get("product").URL("category", "electronics", "id", "123")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Built Product URL:", productURL.String()) // Output: "/products/electronics/123"
}
