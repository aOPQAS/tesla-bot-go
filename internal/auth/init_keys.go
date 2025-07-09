package auth

import (
	"os"

	"github.com/joho/godotenv"
)

var SigningKeyAccess []byte
var SigningKeyRefresh []byte

func InitJWTKeys() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	accessKey := os.Getenv("ACCESS_SIGNING_KEY")
	refreshKey := os.Getenv("REFRESH_SIGNING_KEY")

	if accessKey == "" || refreshKey == "" {
		panic("JWT signing keys not in .env")
	}

	SigningKeyAccess = []byte(accessKey)
	SigningKeyRefresh = []byte(refreshKey)
}
