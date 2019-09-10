package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"starter-kit/apps"
	"starter-kit/config"
	"starter-kit/constants"
	"starter-kit/utils"
)

func main() {

	//Log ----------------------
	log.SetFlags(0)
	multiWriter := io.MultiWriter(new(utils.LogWriter))
	log.SetOutput(multiWriter)
	//--------------------------
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
	engine.Static(`/public`, `public`)
	apps.RegisterApps(engine)
	err := engine.Run(fmt.Sprintf(`%v:%v`, host, port))
	if err != nil {
		log.Panicf(`Could not start server => %v`, err)
	}

}
