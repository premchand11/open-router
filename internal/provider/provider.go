package provider

import "context"

type ChatRequest struct {
	Model   string    `json:"model"`
	Prompt  string    `json:"prompt"`
}

type ChatResponse struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	TokensUsed int    `json:"tokens_used"`
}

type Provider interface {
	Name() string
	Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error)
	Health(ctx context.Context) error
}