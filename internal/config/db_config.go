package config

import (
	"accounts-and-transactions/internal/logger"
	"context"
	"sync"

	"github.com/Netflix/go-env"
)

type DBConfig struct {
	Host                string `env:"DB_HOST"`
	Port                string `env:"DB_PORT"`
	UserName            string `env:"DB_USER_NAME"`
	Password            string `env:"DB_PASSWORD"`
	DatabaseName        string `env:"DB_NAME"`
	AutoMigrateRequired bool   `env:"AUTO_MIGRATE_REQUIRED,default:false"`
}

var (
	onceDBConfig sync.Once
	dbConfig     DBConfig
)

func GetDBConfig(ctx context.Context) *DBConfig {
	onceDBConfig.Do(func() {
		_, err := env.UnmarshalFromEnviron(&dbConfig)
		if err != nil {
			logger.WithContext(ctx).Errorf("error unmarshalling db server config, err: %v", err)
			panic(err)
		}
	})
	return &dbConfig
}
