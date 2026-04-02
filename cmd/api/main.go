package main

import (
	"log"
	"net/http"
	"os"

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
	db := database.NewPostgresDB()
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
	r.Post("/api/v1/auth/refresh", authHandler.Refresh)
	r.Post("/api/v1/auth/logout", authHandler.Logout)

	r.Group(func(admin chi.Router) {
		admin.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		admin.Use(middleware.RequiredRole("admin"))
		admin.Get("/api/v1/admin", userHandler.AdminOnly)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("server is running on port" + cfg.Port))
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
		<html>
		<head>
			<title>Auth Service</title>
			<style>
				body {
					font-family: Arial;
					text-align: center;
					margin-top: 50px;
				}
				h1 { color: #4CAF50; }
				ul { list-style: none; }
			</style>
		</head>
		<body>
			<h1>🚀 Auth Service</h1>
			<p>Backend API is running successfully</p>

			<h3>Endpoints:</h3>
			<ul>
				<li>/health</li>
				<li>/api/v1/auth/login</li>
				<li>/api/v1/auth/refresh</li>
				<li>/api/v1/users/me</li>
				<li>/api/v1/admin</li>
			</ul>

			<p>Made by Chintu Kumar 💻</p>
		</body>
		</html>
	`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("server runnnig on port", port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("server failed", err)
	}
}
