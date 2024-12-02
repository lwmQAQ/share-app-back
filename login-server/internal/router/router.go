package router

import (
	"login-server/internal/handle/userhandler"
	"login-server/internal/svc"
	"login-server/middleware"

	"github.com/gin-gonic/gin"
)

// AddRouter 添加路由
func AddRouter(r *gin.Engine, svc *svc.ServiceContext) *gin.Engine {
	// 创建 v1 路由组（使用中间件进行用户验证）
	v1 := r.Group("/v1/api", middleware.UserVerifyMiddleware(svc.JWTUtil))
	{
		v1.GET("/userinfo", func(c *gin.Context) {
			userhandler.GetUserInfoHandler(c, svc)
		})
		v1.PUT("/update", func(c *gin.Context) {
			userhandler.UpdateUserHandler(c, svc)
		})
	}
	// 创建 v1/public 路由组（不使用中间件）
	v2 := r.Group("/v1/api/public")
	{
		v2.POST("/login", func(c *gin.Context) {
			userhandler.LoginHandler(c, svc)
		})
		v2.POST("/register", func(c *gin.Context) {
			userhandler.RegisterHandler(c, svc)
		})
		v2.GET("/get/code", func(c *gin.Context) {
			userhandler.SendCodeHandler(c, svc)
		})
		v2.POST("/login/code", func(c *gin.Context) {
			userhandler.LoginByCoderHandler(c, svc)
		})
	}

	return r // 返回修改后的路由
}
