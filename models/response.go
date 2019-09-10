package models

import (
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

func NewResponse(status int, message string, payload interface{}) *Response {
	res := Response{
		Status:    status,
		Message:   message,
		Payload:   payload,
		Timestamp: time.Now(),
	}
	return &res
}
