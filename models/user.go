package models

import (
	"github.com/DowneyL/now/packages/util"
	"log"
	"time"
)

type User struct {
	BaseModel
	Name         string    `gorm:"unique;not null;size:50;default:'';comment:'用户名'" json:"name"`
	Password     string    `gorm:"not null;size:100;default:'';comment:'密码'" json:"-"`
	Email        string    `gorm:"index;not null;size:120;default:'';comment:'用户邮箱'" json:"email"`
	State        *int      `gorm:"not null;type:tinyint(2);default:'2';comment:'状态, 1: 启用 2: 未认证 -1: 禁用'" json:"state"`
	VerifiedTime time.Time `gorm:"not null;default:'1970-01-01 00:00:00';comment:'认证时间'" json:"verified_time"`
}

func FindUserByName(name string) (user User) {
	db.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(name string, password string) *User {
	user := &User{}
	user.Name = name
	user.Password = util.MustGeneratePassword(password)
	if err := wdb.Create(user).Error; err != nil {
		log.Fatalln(err)
		return nil
	}

	return user
}
