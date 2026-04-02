package model

import "time"

type refreshToken struct{
	ID string
	UserID string
	Token string
	ExpiresAt time.Time
	CreatedAt time.Time
}



