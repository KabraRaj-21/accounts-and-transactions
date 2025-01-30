package router

import (
	"context"
	"transaction/internal/server/http/handler"
	"transaction/internal/server/http/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(ctx context.Context, accountHandler *handler.AccountHandler, transactionHandler *handler.TransactionHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware())

	setupAccountRoutes(router, accountHandler)
	setupTransactionRoutes(router, transactionHandler)

	return router
}

func setupAccountRoutes(router *gin.Engine, accountHandler *handler.AccountHandler) {
	accountRoutes := router.Group("/api/v1/accounts")
	{
		accountRoutes.POST("/", accountHandler.CreateAccount)
		accountRoutes.GET("/:accountId", accountHandler.GetAccount)
	}
}

func setupTransactionRoutes(router *gin.Engine, transactionHandler *handler.TransactionHandler) {
	accountRoutes := router.Group("/api/v1/transactions")
	{
		accountRoutes.POST("/", transactionHandler.RegisterTransaction)
	}
}
