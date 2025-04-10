package libmiddleware

import "github.com/golang-jwt/jwt/v4"

type JWTGenerator interface {
	Generate(userID string) (string, error)
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
