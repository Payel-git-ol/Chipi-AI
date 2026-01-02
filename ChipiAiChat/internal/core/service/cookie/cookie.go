package cookie

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
)

func GetJWT(c echo.Context) (string, error) {
	if cookie, err := c.Cookie("jwt"); err == nil {
		return cookie.Value, nil
	}

	auth := c.Request().Header.Get("Authorization")
	if auth != "" {
		return auth, nil
	}

	if token := c.QueryParam("token"); token != "" {
		return token, nil
	}

	return "", fmt.Errorf("no token provided")
}

func ValidateJWT(tokenStr string, secret []byte) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func CheckJWT(c echo.Context) (string, error) {
	token, err := GetJWT(c)
	if err != nil {
		return "", err
	}

	log.Println("token:", token)

	return token, nil
}
