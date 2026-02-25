package repository

import "github.com/premchand11/open-router/internal/server"

type Repositories struct {
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
