package provider

import (
	"context"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type MockProvider struct{}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (m *MockProvider) Name() string {
	return "mock"
}

func (m *MockProvider) Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error) {
	// Simulate latency
	time.Sleep(time.Duration(rand.Intn(200)+50) * time.Millisecond)

	return &ChatResponse{
		ID:         uuid.New().String(),
		Content:    "Mock response to: " + req.Prompt,
		TokensUsed: rand.Intn(100) + 10,
	}, nil
}

func (m *MockProvider) Health(ctx context.Context) error {
	// Always healthy for now
	return nil
}
