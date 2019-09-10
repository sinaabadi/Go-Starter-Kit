package apiV1

import (
	"github.com/gin-gonic/gin"
	"starter-kit/models"
	"starter-kit/utils"
)

func RegisterIndexRoutes(router *gin.RouterGroup, deps ...interface{}) *gin.RouterGroup {
	router.GET(`/`, func(context *gin.Context) {
		utils.FormatAndSend(context, models.Response{
			Status:  200,
			Message: "Hi",
			Payload: nil,
		})
	})

	return router
}
