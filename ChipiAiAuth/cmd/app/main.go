package main

import (
	"ChipiAiAuth/internal/fetcher/kafka/consumer"
	"ChipiAiAuth/pkg/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func main() {
	database.InitDb()
	e := echo.New()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.GetMessageNewUser(&wg)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":7070"))
}
