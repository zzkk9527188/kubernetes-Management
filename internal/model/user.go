package model

import (
	"gorm.io/gorm"
	"time"
)

// User 定义用户表的结构
type User struct {
	ID       uint           `gorm:"primary_key"`
	Username string         `gorm:"unique;not null" json:"username"`
	Password string         `gorm:"not null" json:"password"`
	Email    string         `gorm:"email" json:"email"`
	CreateAt time.Time      `gorm:"autoCreateTime" json:"createAt"`
	UpdateAt time.Time      `gorm:"autoUpdateTime" json:"updateAt"`
	DeleteAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}
