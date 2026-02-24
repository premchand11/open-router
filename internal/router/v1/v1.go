package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/premchand11/or/internal/handler"
	"github.com/premchand11/or/internal/middleware"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middleware *middleware.Middlewares) {
	// Register chat routes
	registerChatRoutes(router, handlers.Chat, middleware.Auth)
}
