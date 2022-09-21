package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//User ..
type User struct {
	UserID    uuid.UUID  `json:"user_id"`
	FirstName string     `json:"first_name" validate:"required"`
	LastName  string     `json:"last_name" validate:"required"`
	Email     string     `json:"email" validate:"required,email,lte=320"`
	Password  string     `json:"password" validate:"required,min=6,max=18"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
