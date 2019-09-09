package apps

import (
	"github.com/gin-gonic/gin"
	"starter-kit/apps/api"
)

func RegisterApps(engine *gin.Engine) {
	registerApiV1Routes(engine)
}

func registerApiV1Routes(engine *gin.Engine) {
	apiV1Group := engine.Group(`/api/v1`)
	api.RegisterApiV1(apiV1Group)
}
