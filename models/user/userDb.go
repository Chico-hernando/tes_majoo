package user

import "time"

type User struct {
	Id    int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
