package api

import (
	"github.com/gin-gonic/gin"
	"starter-kit/middlewares"
	apiV1 "starter-kit/routes/api/v1"
)

func RegisterApiV1(app *gin.RouterGroup) {

	mongoClient := middlewares.ConnectToMongo()

	var v1Routes = map[string]func(router *gin.RouterGroup, deps ...interface{}) *gin.RouterGroup{
		`/`: apiV1.RegisterIndexRoutes,
	}

	for path, routeFunction := range v1Routes {
		indexGroup := app.Group(path)
		routeFunction(indexGroup)
	}

}
