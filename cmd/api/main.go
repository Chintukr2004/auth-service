package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running"))
	})

	port := ":8080"
	log.Println("server runnnig on port", port)

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("server failed", err)
	}
}
