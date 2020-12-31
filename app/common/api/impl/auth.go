package impl

import (
	"gin-mvc/app/common/service"
	"gin-mvc/app/common/util/ApiUtil"
	"github.com/gin-gonic/gin"
	"strings"
)

type loginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	form := &loginForm{}
	if c.BindJSON(&form) != nil {
		ApiUtil.JSON(nil, "参数错误", ApiUtil.FailCode, c)
	} else if strings.Trim(form.Username, " ") == "" || strings.Trim(form.Password, " ") == "" {
		ApiUtil.JSON(nil, "用户名或密码为空", ApiUtil.FailCode, c)
	} else {
		isOk, info := service.LoginVerify(form.Username, form.Password)
		if !isOk {
			ApiUtil.JSON(nil, "用户名或密码错误", ApiUtil.FailCode, c)
		} else {
			ApiUtil.JSON(info, "登录成功", ApiUtil.SuccessCode, c)
		}
	}
}

func Logout(c *gin.Context) {
	uid := c.GetInt64("uid")
	if service.Logout(uid) {
		ApiUtil.JSON(nil, "退出成功", ApiUtil.SuccessCode, c)
	} else {
		ApiUtil.JSON(nil, "退出失败", ApiUtil.FailCode, c)
	}
}
