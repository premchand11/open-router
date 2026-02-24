package service

import (
	"github.com/premchand11/OR/internal/provider"
	"github.com/premchand11/OR/internal/repository"
	"github.com/premchand11/OR/internal/server"
)

type Services struct {
	Chat *ChatService
}

func NewServices(s *server.Server, repos *repository.Repositories, registry *provider.Registry) (*Services, error) {
	return &Services{
		Chat: NewChatService(registry),
	}, nil
}
