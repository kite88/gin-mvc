package api

import (
	"gin-mvc/app/common/util/ApiUtil"
	"gin-mvc/app/user/service"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	ApiUtil.JSON(service.QueryUsers(), "获取数据成功", ApiUtil.SuccessCode, c)
}
