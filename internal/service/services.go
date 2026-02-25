package service

import (
	"github.com/premchand11/open-router/internal/provider"
	"github.com/premchand11/open-router/internal/repository"
	"github.com/premchand11/open-router/internal/server"
)

type Services struct {
	Chat *ChatService
}

func NewServices(s *server.Server, repos *repository.Repositories, registry *provider.Registry) (*Services, error) {
	return &Services{
		Chat: NewChatService(registry),
	}, nil
}
