package apiV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterIndexRoutes(router *gin.RouterGroup, deps ...interface{}) *gin.RouterGroup {
	router.GET(`/`, func(context *gin.Context) {
		context.JSON(http.StatusOK, struct {
			Status  int `json:"status"`
			Message string `json:"message"`
		}{
			Status:  200,
			Message: "Hi ಠ_ಠ Api V1 running smoothly",
		})
	})

	return router
}
