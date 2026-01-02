package chat

import "github.com/labstack/echo/v4"

type ChatRepository interface {
	ChatInAi(c echo.Context) error
}
