package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", http.FileServer(http.Dir("./static")).ServeHTTP)
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Println(err)
	}
}
