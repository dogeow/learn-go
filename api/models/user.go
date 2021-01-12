package models

import (
	"time"
)

// gorm.Model 的定义
type Model struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Model
	Name     string
	Email    string
	Password string
}

type Login struct {
	Name     string `form:"Name" json:"Name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
