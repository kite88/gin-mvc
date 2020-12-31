package model

import (
	"gin-mvc/app/common/model"
	"gin-mvc/conf"
)

type UserModel struct {
	model.Model
	Username string `json:"username"`
	Password string `json:"-"`
}

func (UserModel) TableName() string {
	return conf.DbTablePrefix + "user"
}
