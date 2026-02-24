package handler

import (
	"github.com/premchand11/or/internal/server"
	"github.com/premchand11/or/internal/service"
)

type Handlers struct {
	Health *HealthHandler
	Chat   *ChatHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		Health: NewHealthHandler(s),
		Chat:   NewChatHandler(services.Chat),
	}

}
