package services

import (
	"github.com/DowneyL/now/models"
	"github.com/DowneyL/now/packages/util"
)

func Register(username, password string) (user *models.User, auth util.Auth, err error) {
	tx := models.WriteDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err = models.CreateUser(username, password, tx)
	if err != nil {
		tx.Rollback()
		return
	}

	auth, err = util.GenerateAuth(username, password)
	if err != nil {
		tx.Rollback()
		return
	}

	return user, auth, tx.Commit().Error
}
