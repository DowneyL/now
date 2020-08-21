package models

import (
	"github.com/DowneyL/now/pkg/util"
)

func Migrate() error {
	// Migrate table `user`
	rootName := "august5th"
	user := &User{}
	if err := WriteDB.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	WriteDB.Where("name = ?", rootName).First(user)
	if user.ID == 0 {
		user.Name = rootName
		user.Password = util.MustGeneratePassword("123456")
		if err := WriteDB.Create(user).Error; err != nil {
			return err
		}
	}

	// Migrate table `occupation`
	if err := WriteDB.AutoMigrate(&Occupation{}).Error; err != nil {
		return err
	}

	return nil
}
