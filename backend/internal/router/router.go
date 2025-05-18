package router

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "cashback_info/cmd/docs"
	tokenservice "cashback_info/internal/service/token"
)

func SetupRouter(tokenService tokenservice.TokenService) *gin.Engine {
	router := gin.Default()

	router.Use(AuthMiddleware(tokenService))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

func AuthMiddleware(tokenService tokenservice.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		var userID uuid.UUID

		if authHeader == "" {
			c.Next()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		var err error
		userID, err = tokenService.ValidateToken(token)
		if err != nil {
			// Если токен недействителен, продолжаем выполнение, но userID будет пустым
			c.Next()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
