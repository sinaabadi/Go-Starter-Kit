package models

import (
	"time"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Payload interface{} `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
