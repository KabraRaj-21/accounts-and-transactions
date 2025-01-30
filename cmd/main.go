package cmd

import (
	"accounts-and-transactions/internal/app"
	"accounts-and-transactions/internal/config"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appConfig := config.GetAppConfig(ctx)
	application := app.New(ctx, appConfig)

	application.Run(ctx)
}
