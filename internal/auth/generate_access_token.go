package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

const accesstokenTTL = 2 * time.Hour

func InitAccessSigningKey() error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	AccessKey := os.Getenv("ACCESS_SIGNING_KEY")
	if AccessKey == "" {
		panic("REFRESH_SIGNING_KEY is not set in .env file")
	}
	SigningKeyAccess = []byte(AccessKey)
	return nil
}

func GenerateAccessToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(accesstokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKeyAccess)
}
