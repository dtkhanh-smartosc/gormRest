package entities

import (
	"time"
)

type User struct {
	Id          uint64    `json:"id" gorm:"primaryKey,autoIncrement"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email" gorm:"unique"`
	PhoneNumber string    `json:"phone_number" gorm:"unique"`
	Password    string    `json:"password"`
	IsVerify    bool      `json:"is_verify" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
