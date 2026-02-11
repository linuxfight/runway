{{- if .Options.Infra.Redis }}
package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"

	"{{ .ModulePath }}/app/config"
)

func NewRedis(cfg *config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}

func RegisterLifecycle(lc fx.Lifecycle, rdb *redis.Client) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return rdb.Close()
		},
	})
}
{{- end }}