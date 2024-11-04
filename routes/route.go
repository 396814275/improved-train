package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web2/controllers"
	"web2/logger"
	"web2/middlewares"
)

func Setup(mode string) *gin.Engine {

	//if mode == "release" {
	//	gin.SetMode(gin.ReleaseMode) //gin设置成发布模式
	//}

	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.LoadHTMLFiles("dist/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := r.Group("/api/v1")
	//注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用jwt认证中间件

	//v1.GET("/", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	//如果用户已经登录,判断请求头中是否有 有效的token
	//	c.JSON(200, settings.Conf.Version)
	//})

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts/", controllers.GetPostListHandler)

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
