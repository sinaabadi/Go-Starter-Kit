package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"starter-kit/config"
)

func main() {
	config := config.GetConfig()
	engine := gin.Default()
	host := config.Get(`host`).(string)
	port := config.Get(`port`).(string)
	goEnv :=
	engine.Run(fmt.Sprintf(`%v:%v`, host, port))

}
