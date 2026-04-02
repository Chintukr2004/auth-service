package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Chintukr2004/auth-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg *config.Config) *pgxpool.Pool {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=require",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Println("Unable to connect to database:", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}
	log.Println("Connected to postgreSql")
	return db
}
