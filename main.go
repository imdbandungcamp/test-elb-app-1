package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/health", func(r chi.Router) {
		r.Get("/", health())
	})
	fmt.Println("Run APP1")
	http.ListenAndServe(":8080", r)
}

func health() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "APP 1: 4")
	})
}
