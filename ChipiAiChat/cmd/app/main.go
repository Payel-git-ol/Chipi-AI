package main

import (
	"ChipiAiChat/internal/core/service/chat"
	"ChipiAiChat/internal/fetcher/grpc/callback"
	"ChipiAiChat/internal/fetcher/grpc/callbackpb"
	"ChipiAiChat/internal/fetcher/kafka/consumer"
	"ChipiAiChat/pkg/database"
	"ChipiAiChat/pkg/models/request"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ws := chat.Chat{}
	e := echo.New()

	database.ConnectMongo()

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

	e.POST("/create/room", func(c echo.Context) error {
		var req request.Room
		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		id, err := database.CreateNewRoom(req.Name)
		if err != nil {
			return c.JSON(500, err.Error())
		}

		return c.JSON(200, map[string]string{
			"roomId": id.Hex(),
		})
	})

	e.GET("/chat", ws.ChatInAi)

	e.Logger.Fatal(e.Start(":7070"))
}
