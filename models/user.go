package models

import (
	"time"
)

type User struct {
	Cellphone string      `json:"cellphone"`
	Username  string      `json:"username"`
	FirstName  interface{} `json:"first_name"`
	LastName  interface{} `json:"last_name"`
	CreatedAt time.Time   `json:"timestamp"`
}
