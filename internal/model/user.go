package model

import (
	"test-agit/internal/abstraction"
)

type User struct {
	Name    string `gorm:"type:varchar(100);not null;" json:"name" validate:"required,max=100"`
	Address string `gorm:"type:varchar(225);not null;" json:"address" validate:"required,max=225"`
}

type UserModel struct {
	abstraction.EntityAI
	User
	abstraction.Filter
}

func (UserModel) TableName() string {
	return "users"
}
