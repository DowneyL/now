package models

import (
	"time"
)

type User struct {
	BaseModel
	Name         string    `gorm:"unique;not null;size:50;default:'';comment:'用户名'" json:"name"`
	Password     string    `gorm:"not null;size:100;default:'';comment:'密码'" json:"-"`
	Email        string    `gorm:"unique;not null;size:120;default:'';comment:'用户邮箱'" json:"email"`
	State        *int      `gorm:"not null;type:tinyint(2);default:'2';comment:'状态, 1: 启用 2: 未认证 -1: 禁用'" json:"state"`
	VerifiedTime time.Time `gorm:"not null;default:'1970-01-01 00:00:00';comment:'认证时间'" json:"verified_time"`
}
