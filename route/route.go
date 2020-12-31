package route

import (
	"fmt"
	commonApi "gin-mvc/app/common/api/impl"
	"gin-mvc/app/common/api/middleware"
	userApi "gin-mvc/app/user/api"
	"gin-mvc/conf"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Static(r *gin.Engine) {
	fmt.Println("加载static静态资源", time.Now().Format(conf.TimeFormatLayout))
	r.StaticFS("/static", http.Dir("static"))
}

func Web(r *gin.Engine) {
	fmt.Println("加载web页面", time.Now().Format(conf.TimeFormatLayout))
	//自定义分隔符 因为现在都流行前后端分离，前端来渲染。避免与前端一些框架冲突，例如vue框架
	r.Delims("{[{", "}]}")
	//加载静态目录文件
	r.LoadHTMLGlob("web/*")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
}

func Api(r *gin.Engine) {
	fmt.Println("加载api路由", time.Now().Format(conf.TimeFormatLayout))
	api := r.Group("/api")
	{
		api.POST("/login", commonApi.Login)
		authorized := api.Group("")
		authorized.Use(middleware.AuthMiddleware)
		{
			authorized.POST("/logout", commonApi.Logout)
			authorized.GET("/user", userApi.List)
		}
	}
}
