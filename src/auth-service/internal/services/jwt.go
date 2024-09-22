package services

import (
	"os"
	"time"

	"auth-service/pkg"
	"github.com/golang-jwt/jwt/v5"
)

// CreateJWT generates a JWT for the provided identity and email.
func CreateJWT(id uint, username, email, role string) (string, error) {
    issuer := os.Getenv("JWT_ISSUER")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   id,
        "username": username,
        "email":    email,
        "role":     role,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
        "iss":      issuer,
    })

    tokenString, err := token.SignedString(pkg.PrivateKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}