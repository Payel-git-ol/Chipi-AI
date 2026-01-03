package main

import (
	"ChipiAiChat/internal/core/service/chat"
	"ChipiAiChat/internal/fetcher/grpc/callback"
	"ChipiAiChat/internal/fetcher/grpc/callbackpb"
	"ChipiAiChat/internal/fetcher/kafka/consumer"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ws := chat.Chat{}
	e := echo.New()

	go func() {
		lis, err := net.Listen("tcp", ":50053")
		if err != nil {
			log.Printf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		callbackpb.RegisterAiCallbackServer(s, &callback.CallbackServer{})

		log.Println("ContextEnhancementService :50053")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go consumer.GetMessageUser(&wg)

	e.GET("/chat", ws.ChatInAi)

	e.Logger.Fatal(e.Start(":7070"))
}
