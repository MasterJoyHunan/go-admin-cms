package routers

import (
	"blog/controller"
	_ "blog/docs"
	"blog/middleware"
	"blog/pkg/setting"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.ApplicationConf.Env)
	r := gin.New()

	// swagger 文档输出
	r.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// vue build 文件目录
	r.LoadHTMLFiles("./dist/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")
	r.Static("/static", "./dist/static")

	// 加入通用中间件
	r.Use(
		gin.Recovery(),           // recovery 防止程序奔溃
		middleware.Logger(),      // 日志记录
		middleware.NoFound(),     // 404 处理
		middleware.MakeSession(), // session支持
		middleware.ErrorHandle(), // 错误处理
	)

	r.GET("/captcha", controller.Captcha)
	r.POST("/login", controller.Login)

	// 加入鉴权中间件
	r.Use(middleware.JWT(), middleware.Authentication())

	// 用户
	{
		r.GET("/auth/role", controller.RoleTree)
		r.GET("/adminUser", controller.UserIndex)
		r.GET("/adminUser/:id", controller.UserDetail)
		r.POST("/adminUser", controller.UserAdd)
		r.PUT("/adminUser/:id", controller.UserEdit)
		r.DELETE("/adminUser/:id", controller.UserDel)
	}

	// 角色
	{
		r.GET("/auth/tree", controller.AuthTree)
		r.GET("/role", controller.RoleIndex)
		r.GET("/role/:id", controller.RoleDetail)
		r.POST("/role", controller.RoleAdd)
		r.PUT("/role/:id", controller.RoleEdit)
		r.DELETE("/role/:id", controller.RoleDel)
	}
	return r
}
