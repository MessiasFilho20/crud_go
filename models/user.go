package models

import "time"

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   uint8  `json:"age"`
	CPF   string `json:"CPF" gorm:"unique"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
