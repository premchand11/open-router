package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/premchand11/open-router/internal/handler"
	"github.com/premchand11/open-router/internal/middleware"
)

func registerChatRoutes(r *echo.Group, h *handler.ChatHandler, auth *middleware.AuthMiddleware) {

	//chat operations
	chat := r.Group("/chat")
	chat.Use(auth.RequireAuth)

	//chat collerctions
	chat.POST("/completions", h.Chat)

}
