package db

import (
	"fmt"
	"ml_daily_record/pkg/models"
	"ml_daily_record/pkg/util"
)

type UserDao interface {
	FindUserByUsernameAndPassword(string, string) (*models.User, error)
	UpdateUserToken(*models.User) error
}

func (pg *PGService) FindUserByUsernameAndPassword(uname string, passwd string) (*models.User, error) {
	sql := fmt.Sprintf("username = '%s' and password = '%s'", uname, util.EncodeToSHA256(passwd))
	userQuery := &models.User{}
	if err := pg.Connection.Where(sql).First(&userQuery); err.Error != nil {
		if err.RecordNotFound() {
			return nil, nil
		} else {
			return nil, err.Error
		}
	}
	if userQuery.Id == "" {
		return nil, nil
	}
	return userQuery, nil
}

func (pg *PGService) UpdateUserToken(user *models.User) error {
	if err := pg.Connection.Model(&user).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"token_uid":         user.TokenUid,
		"token":             user.Token,
		"token_create_time": user.TokenCreateTime,
		"last_login_time":   user.LastLoginTime,
	}).Error; err != nil {
		return err
	}
	return nil
}
