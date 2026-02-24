package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/premchand11/or/internal/provider"
	"github.com/premchand11/or/internal/service"
)

type ChatHandler struct {
	chatService *service.ChatService
}

func NewChatHandler(chatService *service.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

func (h *ChatHandler) Chat(c echo.Context) error {
	var req provider.ChatRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	resp, err := h.chatService.Chat(c.Request().Context(), req)
	if err != nil {
		if errors.Is(err, provider.ErrProviderNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "invalid model",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, resp)
}
