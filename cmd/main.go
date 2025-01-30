package cmd

import (
	"context"
	"transaction/internal/app"
	"transaction/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	appConfig := config.GetAppConfig(ctx)
	application := app.New(ctx, appConfig)

	application.Run(ctx)
}
