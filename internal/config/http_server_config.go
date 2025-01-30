package config

import (
	"accounts-and-transactions/internal/logger"
	"context"
	"sync"

	"github.com/Netflix/go-env"
)

type HttpServerConfig struct {
	Port string `env:"HTTP_PORT,default=8080"`
}

var (
	onceHttpServerConfig sync.Once
	httpServerConfig     HttpServerConfig
)

func GetHTTPServerConfig(ctx context.Context) *HttpServerConfig {
	onceHttpServerConfig.Do(func() {
		_, err := env.UnmarshalFromEnviron(&httpServerConfig)
		if err != nil {
			logger.WithContext(ctx).Errorf("error unmarshalling http server config, err: %v", err)
			panic(err)
		}
	})
	return &httpServerConfig
}
