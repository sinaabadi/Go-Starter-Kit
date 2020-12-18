package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"log"
	"starter-kit/apps"
	"starter-kit/config"
	"starter-kit/constants"
	"starter-kit/middlewares"
	"starter-kit/utils"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title Golang Starter kit API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

func main() {

	//Log ----------------------
	log.SetFlags(0)
	multiWriter := io.MultiWriter(new(utils.LogWriter))
	log.SetOutput(multiWriter)
	appConfig := config.GetConfig()
	appEnv := appConfig.Get(`appEnv`).(string)
	if appEnv == constants.PRODUCTION_MODE {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	host := appConfig.Get(`host`).(string)
	port := appConfig.Get(`port`).(string)
	utils.InitI18n()
	engine.Static(`/public`, `public`)
	apps.RegisterApps(engine)
	engine.NoRoute(middlewares.NoRouteHandler)
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := engine.Run(fmt.Sprintf(`%v:%v`, host, port))
	if err != nil {
		log.Panicf(`Could not start server => %v`, err)
	}

}
