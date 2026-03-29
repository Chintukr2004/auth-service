package main

import (
	"log"
	"net/http"

	"github.com/Chintukr2004/auth-service/internal/config"
	"github.com/go-chi/chi/v5"
)

func main() {

	cfg := config.LoadConfig()

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running on port" + cfg.Port))
	})

	log.Println("server runnnig on port", cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		log.Fatal("server failed", err)
	}
}
