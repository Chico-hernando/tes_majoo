package user

import (
	"time"
)

type UserResponse struct {
	Id        int    `json:"id"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email" gorm:"unique"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
