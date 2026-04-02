package repository

import (
	"context"
	"time"

	"github.com/Chintukr2004/auth-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, role, is_verified, created_at, updated_at
	`

	return r.db.QueryRow(ctx, query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.Role,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, name, email, password_hash, role, is_verified, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user model.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	query := `
		SELECT id, name, email, password_hash, role, is_verified, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) SaveRefreshToken(ctx context.Context, userID, token string, expiresAt time.Time) error {
	query := `
		INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
		`
	_, err := r.db.Exec(ctx, query, userID, token, expiresAt)
	return err
}

func (r *UserRepository) GetUserByRefreshToken(ctx context.Context, token string) (string, error) {
	query := `
		SELECT user_id FROM refresh_tokens
		WHERE token = $1 AND expires_at > NOW()
		`

	var userID string
	err := r.db.QueryRow(ctx, query, token).Scan(&userID)
	return userID, err

}

func (r *UserRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	query := `
		DELETE FROM refresh_tokens WHERE token = $1
		`
	_, err := r.db.Exec(ctx, query, token)
	return err
}
