package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"starter-kit/models"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func FormatAndSend(context *gin.Context, res *models.Response) {
	responseBytes, err := json.Marshal(res)
	if err != nil {
		log.Printf(`Unable to marshal response => %v`, err)
	}
	context.JSON(res.Status, res)
	log.Println(string(responseBytes))
}
