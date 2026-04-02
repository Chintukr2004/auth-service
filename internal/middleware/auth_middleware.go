package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			//format:bearer <token>

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			userID, ok := claims["user_id"].(string)
			if !ok {
				http.Error(w, "Invalid user ID", http.StatusUnauthorized)
				return
			}
			role,_ := claims["role"].(string)

			//Add userID to context
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, "role", role)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func RequiredRole(requiredRole string)func(http.Handler)http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			
			role, ok:= r.Context().Value("role").(string)
			if !ok || role != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return 
			}
			next.ServeHTTP(w,r)
		})
	}
}