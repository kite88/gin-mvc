package dao

import (
	"gin-mvc/app/common/dao"
	"gin-mvc/app/user/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
}

var (
	m    model.UserModel
	list []model.UserModel
)

func (this UserDao) Dao() *gorm.DB {
	return dao.DB().Table(m.TableName())
}

func (this UserDao) Query(username string) (bool, model.UserModel) {
	r := this.Dao().Where("username = ?", username).First(&m)
	if r.RowsAffected == 0 || r.Error != nil {
		return false, m
	}
	return true, m
}

func (this UserDao) Add(userModel model.UserModel) bool {
	r := this.Dao().Create(&userModel)
	if r.Error != nil || r.RowsAffected == 0 {
		return false
	}
	return true
}

func (this UserDao) QueryList() []model.UserModel {
	this.Dao().Find(&list)
	return list
}
