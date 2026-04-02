package service

import (
	"context"
	"errors"
	"time"

	"github.com/Chintukr2004/auth-service/internal/model"
	"github.com/Chintukr2004/auth-service/internal/repository"
	"github.com/Chintukr2004/auth-service/internal/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (*model.User, error) {

	// check if user already exists
	existingUser, _ := s.userRepo.GetUserByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("emial already exists")
	}

	//hash password
	hashedpassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedpassword,
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, password string, jwtSecret string) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", errors.New("Invalid credentials")
	}
	//check password
	err = utils.CheckPassword(password, user.PasswordHash)
	if err != nil {
		return "", "", errors.New("Invalid credentials")
	}

	//Generate tokens
	accessToken, err := utils.GenerateToken(user.ID, jwtSecret, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateToken(user.ID, jwtSecret, 7*24*time.Hour)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)
	err = s.userRepo.SaveRefreshToken(ctx, user.ID, refreshToken, expiresAt)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
func (s *AuthService) Refresh(ctx context.Context, refreshToken, jwtSecret string) (string, error) {
	userID, err := s.userRepo.GetUserByRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}
	newAccessToken, err := utils.GenerateToken(userID, jwtSecret, 15*time.Minute)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}


func(s *AuthService)Logout(ctx context.Context, refreshToken string) error{
	return s.userRepo.DeleteRefreshToken(ctx, refreshToken)
}