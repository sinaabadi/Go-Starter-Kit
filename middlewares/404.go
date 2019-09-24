package middlewares

import (
	"github.com/gin-gonic/gin"
	"starter-kit/utils"
)

func NoRouteHandler(context *gin.Context) {
	utils.FormatAndSend(context, 404, `Resource not found`, nil	)
}
