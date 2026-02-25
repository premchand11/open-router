package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/premchand11/open-router/internal/handler"
	"github.com/premchand11/open-router/internal/middleware"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middleware *middleware.Middlewares) {
	// Register chat routes
	registerChatRoutes(router, handlers.Chat, middleware.Auth)
}
