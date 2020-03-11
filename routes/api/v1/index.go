package apiV1

import (
	"github.com/gin-gonic/gin"
	"starter-kit/utils"
)

func RegisterIndexRoutes(router *gin.RouterGroup, deps ...interface{}) *gin.RouterGroup {
	router.GET(`/`,
		func(context *gin.Context) {
			utils.Debug(`this is test`, utils.LogFields{
				"hello":  context,
				"number": 34,
			})
			utils.FormatAndSend(context,
				200,
				utils.Translate("hello", nil, "en"),
				nil)
		})
	return router
}
