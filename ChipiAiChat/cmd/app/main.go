package main

import (
	"ChipiAiChat/internal/core/service/chat"
	"ChipiAiChat/internal/fetcher/kafka/consumer"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ws := chat.Chat{}
	e := echo.New()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.GetMessageUser(&wg)

	e.GET("/chat", ws.ChatInAi)

	e.Logger.Fatal(e.Start(":7070"))
}
