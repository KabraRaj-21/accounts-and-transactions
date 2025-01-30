package config

import (
	"context"
	"sync"
)

type AppConfig struct {
	HttpServerConfig *HttpServerConfig
	DBConfig         *DBConfig
}

var (
	appConfig     AppConfig
	onceAppConfig sync.Once
)

func GetAppConfig(ctx context.Context) *AppConfig {
	onceAppConfig.Do(func() {
		appConfig.DBConfig = GetDBConfig(ctx)
		appConfig.HttpServerConfig = GetHTTPServerConfig(ctx)
	})
	return &appConfig
}
