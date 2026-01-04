package chat

import (
	"ChipiAiChat/internal/core/service/cookie"
	"ChipiAiChat/internal/fetcher/grpc/client"
	"ChipiAiChat/pkg/database"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Chat struct{}

var Connections = map[string]*websocket.Conn{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ch Chat) ChatInAi(c echo.Context) error {
	roomId := c.QueryParam("roomId")

	token, err := cookie.CheckJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "missing token")
	}

	claims, err := cookie.ValidateJWT(token, []byte(os.Getenv("JWT_KEY_CHAT")))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	username := claims["username"].(string)
	fmt.Println("User:", username)

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("Error upgrade:", err)
		return err
	}

	Connections[username] = conn
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error read:", err)
			break
		}

		fmt.Println("User message:", string(message))

		go database.SaveMessage(roomId, username, string(message))
		go client.SendContent(username, string(message), roomId)
	}

	return nil
}
