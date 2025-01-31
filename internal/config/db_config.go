package config

import (
	"accounts-and-transactions/internal/logger"
	"context"
	"sync"

	"github.com/Netflix/go-env"
)

type DBConfig struct {
	Host                string `env:"DB_HOST,default=localhost"`
	Port                string `env:"DB_PORT,default=3306"`
	UserName            string `env:"DB_USER_NAME,default=myuser"`
	Password            string `env:"DB_PASSWORD,default=mypassword"`
	DatabaseName        string `env:"DB_NAME,default=mydb"`
	AutoMigrateRequired bool   `env:"AUTO_MIGRATE_REQUIRED,default=false"`
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
