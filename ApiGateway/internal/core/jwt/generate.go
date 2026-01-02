package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func GenerateJWT(username string, email string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	var key = os.Getenv("JWT_KEY")
	if key == "" {
		return "", errors.New("no JWT_KEY found")
	}
	var jwtKey = []byte(key)

	claims := jwt.MapClaims{
		"username": username,
		"email":    email,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
