package router

import (
	"apps-server/internal/handle"
	"apps-server/internal/svc"

	"github.com/gin-gonic/gin"
)

// AddRouter 添加路由
func AddRouter(r *gin.Engine, svc *svc.ServiceContext) *gin.Engine {
	// 创建 v1 路由组（使用中间件进行用户验证）
	v1 := r.Group("/v1/api")
	{
		v1.GET("/apps/list", func(c *gin.Context) {
			handle.GetAppsList(c, svc)
		})
		v1.GET("/app/detials/:id", func(c *gin.Context) {
			handle.GetAppDetials(c, svc)
		})
	}
	// 创建 v1/public 路由组（不使用中间件）
	return r // 返回修改后的路由
}
