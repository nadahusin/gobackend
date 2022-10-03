package models

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password,omitempty" validate:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Users []User
