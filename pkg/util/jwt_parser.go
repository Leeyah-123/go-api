package util

import (
	"context"
	"fiber/platform/db"
	"fiber/postgres"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	ID uuid.UUID `json:"id"`
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id := claims["id"].(string)

		// parsing id from string to uuid.UUID
		parsedID := uuid.MustParse(id)

		user, err := db.Query().GetUserById(context.Background(), parsedID)
		if err != nil {
			return nil, fmt.Errorf("invalid token provided")
		}
		if user == (postgres.GetUserByIdRow{}) {
			return nil, fmt.Errorf("invalid token provided")
		}

		return &TokenMetadata{
			ID: user.ID,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearerToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}
