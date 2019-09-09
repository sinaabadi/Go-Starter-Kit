package main

import (
	"github.com/gin-gonic/gin"
	"starter-kit/config"
)


func main() {
	config := config.GetConfig()
	engine := gin.Default()


}
