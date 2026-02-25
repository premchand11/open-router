package rate_limit

import (
	"github.com/premchand11/open-router/internal/config"
	"github.com/redis/go-redis/v9"
)

type Limiter struct {
	redis *redis.Client
}

func NewLimiter(cfg *config.Config, rdb *redis.Client) *Limiter {
	return &Limiter{
		redis: rdb,
	}
}
