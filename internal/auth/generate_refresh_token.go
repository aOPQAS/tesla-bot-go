package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

const refreshtokenTTL = 30 * 24 * time.Hour

func InitRefreshSigningKey() error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	RefreshKey := os.Getenv("REFRESH_SIGNING_KEY")
	if RefreshKey == "" {
		panic("REFRESH_SIGNING_KEY is not set in .env file")
	}
	SigningKeyRefresh = []byte(RefreshKey)
	return nil
}

func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshtokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKeyRefresh)
}
