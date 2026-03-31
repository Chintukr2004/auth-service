package main

import (
	"log"
	"net/http"

	"github.com/Chintukr2004/auth-service/internal/config"
	"github.com/Chintukr2004/auth-service/internal/database"
	"github.com/Chintukr2004/auth-service/internal/handler"
	"github.com/Chintukr2004/auth-service/internal/middleware"
	"github.com/Chintukr2004/auth-service/internal/repository"
	"github.com/Chintukr2004/auth-service/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	cfg := config.LoadConfig()

	// db connect
	db := database.NewPostgresDB(cfg)
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService, cfg.JWTSecret)

	userHandler := handler.NewUserHandler()

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		protected.Get("/api/v1/users/me", userHandler.GetProfile)
	})

	r.Post("/api/v1/auth/register", authHandler.Register)
	r.Post("/api/v1/auth/login", authHandler.Login)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running on port" + cfg.Port))
	})

	log.Println("server runnnig on port", cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		log.Fatal("server failed", err)
	}
}
