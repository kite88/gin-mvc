package main

import (
	"fmt"
	"gin-mvc/app/common/api/middleware"
	commonTask "gin-mvc/app/common/task"
	"gin-mvc/conf"
	"gin-mvc/db"
	"gin-mvc/route"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("应用启动")
	gin.SetMode(conf.RunMode)
	r := gin.Default()
	r.Use(middleware.CorsMiddleware) //允许跨域中间件
	route.Static(r)                  //加载静态资源
	route.Web(r)                     //加载web页面
	route.Api(r)                     //加载api路由
	db.MysqlInit()                   //数据库连接
	db.RedisInit()                   //redis连接
	commonTask.Start()               //定时器
	err := r.Run(":" + conf.HttpPort)
	if err != nil {
		log.Fatalln("应用启动失败", err.Error())
	}
}
