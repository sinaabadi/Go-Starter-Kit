package api

import (
	"github.com/gin-gonic/gin"
	apiV1 "starter-kit/routes/api/v1"
)

func RegisterApiV1(app *gin.RouterGroup) {

	//middlewares.ConnectToMongo()

	var v1Routes = map[string]func(router *gin.RouterGroup, deps ...interface{}) *gin.RouterGroup{
		`/`: apiV1.RegisterIndexRoutes,
		`/auth`: apiV1.RegisterAuthRoutes,
	}

	for path, routeFunction := range v1Routes {
		indexGroup := app.Group(path)
		routeFunction(indexGroup)
	}

}
