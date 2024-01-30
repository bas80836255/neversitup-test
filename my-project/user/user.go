package user

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateDate time.Time `json:"create_date" `
	CreateBy   string    `json:"create_by"`
}
