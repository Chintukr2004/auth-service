package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Chintukr2004/auth-service/internal/config"
	"github.com/Chintukr2004/auth-service/internal/database"
	"github.com/Chintukr2004/auth-service/internal/model"
	"github.com/Chintukr2004/auth-service/internal/repository"
	"github.com/go-chi/chi/v5"
)

func main() {

	cfg := config.LoadConfig()

	// db connect
	db := database.NewPostgresDB(cfg)
	defer db.Close()

	repo := repository.NewUserRepository(db)

	user := &model.User{
		Name:         "Test",
		Email:        "test@example.com",
		PasswordHash: "hashedpassword",
	}

	err := repo.CreateUser(context.Background(), user)
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("User created:", user.ID)
	}
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running on port" + cfg.Port))
	})

	log.Println("server runnnig on port", cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		log.Fatal("server failed", err)
	}
}
