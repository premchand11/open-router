package service

import (
	"context"

	"github.com/premchand11/open-router/internal/provider"
)

type ChatService struct {
	registry *provider.Registry
}

func NewChatService(registry *provider.Registry) *ChatService {
	return &ChatService{
		registry: registry,
	}
}

func (s *ChatService) Chat(ctx context.Context, req provider.ChatRequest) (*provider.ChatResponse, error) {
	// For now, always use mock provider
	p, err := s.registry.Get(req.Model)
	if err != nil {
		return nil, err
	}

	return p.Chat(ctx, req)
}
