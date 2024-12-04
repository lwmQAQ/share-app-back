package router

import (
	"resource-server/internal/handle/posthandle"
	"resource-server/internal/svc"

	"github.com/gin-gonic/gin"
)

// AddRouter 添加路由
func AddRouter(r *gin.Engine, svc *svc.ServiceContext) *gin.Engine {
	// 创建 v1 路由组（使用中间件进行用户验证）
	v1 := r.Group("/v1/api")
	{
		v1.PUT("/create/post", func(c *gin.Context) {
			posthandle.CreatePostHandle(c, svc)
		})
		v1.GET("get/post", func(c *gin.Context) {
			posthandle.SearchDetailsHandle(c, svc)
		})

	}
	return r // 返回修改后的路由
}
