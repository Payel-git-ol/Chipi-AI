package main

import (
	"ApiGateway/internal/core/jwt"
	"ApiGateway/internal/core/logerr"
	"ApiGateway/internal/fetcher/kafka/producer"
	"ApiGateway/pkg/models/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	Logerr := logerr.Logerr{}

	e.POST("/register", func(c echo.Context) error {
		var req request.UserRequest

		if err := c.Bind(&req); err != nil {
			Logerr.LogerrStatusBad("POST", "/register", "200", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		token, err := jwt.GenerateJWT(req.Username, req.Email)
		if err != nil {
			Logerr.LogerrStatusBad("POST", "/register", "200", err)
		}

		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
		}

		c.SetCookie(cookie)

		producer.SendMessageNewUser("register", req)

		Logerr.LogerrStatusOk("POST", "/register", "200")
		return c.JSON(http.StatusOK, token)
	})

	e.POST("/auth", func(c echo.Context) error {
		var req request.UserRequest

		if err := c.Bind(&req); err != nil {
			Logerr.LogerrStatusBad("POST", "/auth", "200", err)
			return c.JSON(http.StatusBadRequest, err)
		}

		producer.SendMessageNewUser("register", req)

		Logerr.LogerrStatusOk("POST", "/auth", "200")
		return c.JSON(http.StatusOK, req)
	})

	//e.POST("create/room", func(c echo.Context) error {

	//})

	e.Logger.Fatal(e.Start(":8080"))
}
