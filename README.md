# Auth Service (Golang)

Production-ready authentication service built using Go.

## Features
- User Signup & Login
- JWT Authentication (Access + Refresh Tokens)
- Password Hashing (bcrypt)
- Protected Routes (Middleware)
- Role-Based Access Control (RBAC)
- Logout (Token Revocation)

## Tech Stack
- Go (Golang)
- PostgreSQL
- pgx
- Chi Router
- JWT

## APIs
- POST /auth/register
- POST /auth/login
- POST /auth/refresh
- POST /auth/logout
- GET /users/me
- GET /admin (admin only)
