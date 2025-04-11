package libmiddleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtGen struct{}

func NewJWTGenerator() JWTGenerator {
	return &jwtGen{}
}

func (g *jwtGen) Generate(userID string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "https://github.com/basputtipong/library",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(rsaPrivateKey)
}
