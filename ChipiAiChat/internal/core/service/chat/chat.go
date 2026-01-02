package chat

import (
	"ChipiAiChat/internal/core/service/cookie"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type Chat struct{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ch Chat) ChatInAi(c echo.Context) error {
	token, err := cookie.CheckJWT(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "missing token")
	}

	claims, err := cookie.ValidateJWT(token, []byte(os.Getenv("JWT_KEY_CHAT")))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "invalid token")
	}

	fmt.Println("User:", claims["username"])

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("Error upgrade:", err)
		return err
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error read:", err)
			break
		}
		fmt.Println(string(message))

		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Error write:", err)
			break
		}
	}
	return nil
}
