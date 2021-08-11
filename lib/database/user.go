package database

import (
	"majoo/configs"
	"majoo/models/user"
)

func GetUser() (dataresult []user.User, err error) {
	err = configs.DB.Find(&dataresult).Error
	if err != nil {
		return nil, err
	}
	return
}