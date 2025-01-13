package controller

import (
	"context"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	cardcontrollers "cashback_info/app/rest_server/controllers/card"
	categorycontrollers "cashback_info/app/rest_server/controllers/category"
	_ "cashback_info/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Serve(ctx context.Context, port int, postgresPool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	router.Use(getDefaultMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerCard(ctx, postgresPool, router)
	registerCategory(ctx, postgresPool, router)

	pprof.Register(router)
	return router
}

func registerCard(ctx context.Context, postgresPool *pgxpool.Pool, router *gin.Engine) {
	cardServer := cardcontrollers.New(ctx, postgresPool)

	router.GET("/users/:user_id/cards", cardServer.ListUserCards)
}

func registerCategory(ctx context.Context, postgresPool *pgxpool.Pool, router *gin.Engine) {
	categoryServer := categorycontrollers.New(ctx, postgresPool)

	router.GET("/categories", categoryServer.ListCategories)
}

func getDefaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle OPTIONS method
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Process the request
		c.Next()
	}
}
