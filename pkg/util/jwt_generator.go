package util

import (
	"fiber/postgres"
	"os"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateToken(user postgres.User) (string, error) {
	// creating jwt claims
	claims := jwt.MapClaims{
		"id": user.ID,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generating encoded token
	encodedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return encodedToken, nil
}
