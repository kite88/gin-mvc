package service

import (
	"gin-mvc/app/common/util/BcryptUtil"
	"gin-mvc/app/user/dao"
	"gin-mvc/app/user/model"
	"time"
)

func CreateUser(username, password string) (int, model.UserModel) {
	userDao := new(dao.UserDao)
	isE, user := userDao.Query(username)
	//判断用户名是否已经存在
	if isE {
		return -2, user
	}
	//构造添加的数据
	var add model.UserModel
	pwdStr, _ := BcryptUtil.Encrypt(password) //密码加密
	add.Username = username
	add.Password = pwdStr
	add.Model.CreateTime = time.Now().Unix()
	//执行添加操作
	res := userDao.Add(add)
	if !res {
		return -1, add
	}
	return 0, add
}

func QueryUsers() []model.UserModel {
	userDao := new(dao.UserDao)
	list := userDao.QueryList()
	return list
}
