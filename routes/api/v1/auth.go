package apiV1

import (
	"github.com/gin-gonic/gin"
	"starter-kit/middlewares"
)

func RegisterAuthRoutes(router *gin.RouterGroup, deps ...interface{}) *gin.RouterGroup {
	authMiddlewares := middlewares.GetAuthMiddlewares()
	router.POST(`/login`, authMiddlewares.LoginHandler)
	router.GET("/refresh-token", authMiddlewares.RefreshHandler)
	router.Use(authMiddlewares.MiddlewareFunc())
	return router
}
