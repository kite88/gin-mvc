package middleware

import (
	"gin-mvc/app/common/service"
	"gin-mvc/app/common/util/ApiUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//解决跨域中间件
func CorsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	c.Next()
}

//登录验证中间件
func AuthMiddleware(c *gin.Context) {
	authorization := c.GetHeader("authorization")
	token := strings.Trim(strings.Trim(authorization, "Bearer"), " ")
	if token == "" {
		ApiUtil.JSON(nil, ApiUtil.UnauthorizedMsg, ApiUtil.UnauthorizedCode, c)
		return
	}
	uid, errMsg := service.TokenVerify(token)
	if uid < 1 {
		ApiUtil.JSON(nil, errMsg, ApiUtil.UnauthorizedCode, c)
		return
	}
	c.Set("uid", uid)
}
