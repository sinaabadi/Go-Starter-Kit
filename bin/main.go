package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"starter-kit/apps"
	"starter-kit/config"
	"starter-kit/constants"
)

func main() {
	appConfig := config.GetConfig()
	appEnv := appConfig.Get(`appEnv`).(string)
	if appEnv == constants.PRODUCTION_MODE {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	host := appConfig.Get(`host`).(string)
	port := appConfig.Get(`port`).(string)
	engine.Static(`/public`, `public`)
	apps.RegisterApps(engine)
	err := engine.Run(fmt.Sprintf(`%v:%v`, host, port))
	if err != nil {
		log.Panicf(`Could not start server => %v`, err)
	}

}
