package models

import (
	"github.com/DowneyL/now/pkg/util"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	BaseModel
	Name         string         `gorm:"unique;not null;size:50;default:'';comment:'用户名'" json:"name"`
	Password     string         `gorm:"not null;size:100;default:'';comment:'密码'" json:"-"`
	Email        string         `gorm:"index;not null;size:120;default:'';comment:'用户邮箱'" json:"email"`
	State        int            `gorm:"not null;type:tinyint(2);default:'2';comment:'状态, 1: 已认证 2: 未认证 -1: 禁用'" json:"state"`
	VerifiedTime *util.DateTime `gorm:"not null;default:'1970-01-01 00:00:00';comment:'认证时间'" json:"verified_time"`
}

const (
	AuthorizedState   = 1
	UnAuthorizedState = 2
	InvalidState      = -1
)

func FindUserByName(name string) (user *User) {
	user = &User{}
	DB.Where("name = ?", name).First(user)
	return
}

func CreateUser(name string, password string, tx *gorm.DB) (user *User, err error) {
	user = &User{}
	db := WriteDB
	if tx != nil {
		db = tx
	}
	user.Name = name
	user.Password = util.MustGeneratePassword(password)
	if err = db.Create(user).Error; err == nil {
		return
	}

	if tx != nil {
		tx.Rollback()
	}
	return
}

func UpdateUserEmail(user *User, email string) error {
	user.Email = email
	user.State = AuthorizedState
	user.VerifiedTime = &util.DateTime{Time: time.Now()}

	return WriteDB.Save(user).Error
}
