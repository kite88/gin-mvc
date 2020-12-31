package service

import (
	"fmt"
	"gin-mvc/app/common/util/ApiUtil"
	"gin-mvc/app/common/util/BcryptUtil"
	"gin-mvc/app/common/util/JwtUtil"
	"gin-mvc/app/user/dao"
	"gin-mvc/conf"
	"gin-mvc/db"
	"strconv"
	"time"
)

var (
	AuthToken string = conf.RedisPrefix + "AUTH_TOKEN_"
	TokenExp  int64  = conf.TokenExp
	JwtSecret string = conf.JwtSecret
)

func LoginVerify(username, password string) (is bool, data map[string]interface{}) {
	userDao := new(dao.UserDao)
	isExist, info := userDao.Query(username)
	if !isExist {
		return
	}
	resBcrypt := BcryptUtil.Decrypt(password, info.Password)
	if !resBcrypt {
		return
	}
	token, exp := JwtUtil.CreateToken(strconv.FormatInt(info.ID, 10), TokenExp, JwtSecret)
	//存入redis
	db.Redis().Set(AuthToken+strconv.FormatInt(info.ID, 10), token, time.Second*time.Duration(TokenExp))
	data = map[string]interface{}{
		"User": info,
		"Token": map[string]interface{}{
			"token": token,
			"exp":   exp,
		},
	}
	is = true
	return
}

func TokenVerify(token string) (uid int64, errMsg string) {
	uidStr, err := JwtUtil.ParseToken(token, conf.JwtSecret)
	uid, _ = strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		errMsg = err.Error()
		return
	}
	//单点登录
	redisToken, redisErr := db.Redis().Get(AuthToken + uidStr).Result()
	if redisErr != nil {
		fmt.Println("redisErr", redisErr.Error())
	}
	if redisToken == "" {
		uid = 0
		errMsg = ApiUtil.UnauthorizedMsg
		return
	}
	if redisToken != token {
		uid = -1
		errMsg = ApiUtil.UnauthorizedSMsg
		return
	}
	return
}

func Logout(uid int64) bool {
	err := db.Redis().Del(AuthToken + strconv.FormatInt(uid, 10)).Err()
	if err != nil {
		return false
	}
	return true
}
