package test

import (
	"fmt"
	commonService "gin-mvc/app/common/service"
	"gin-mvc/app/user/service"
	"gin-mvc/db"
	"strconv"
	"testing"
)

func init() {
	db.MysqlInit()
}

func TestUserAdd(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(service.CreateUser("harden"+strconv.Itoa(i), "123123"))
	}
}

func TestUserVer(t *testing.T) {
	fmt.Println(commonService.LoginVerify("harden", "123123"))
}
