package models

import (
	"github.com/DowneyL/now/packages/util"
)

func Migrate() error {
	rootName := "august5th"
	user := &User{}

	if err := wdb.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	wdb.Where("name = ?", rootName).First(user)
	if user.ID == 0 {
		user.Name = rootName
		user.Password = util.MustGeneratePassword("123456")
		if err := wdb.Create(user).Error; err != nil {
			return err
		}
	}

	return nil
}
