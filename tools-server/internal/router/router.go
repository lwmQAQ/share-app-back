package router

import (
	"tools-back/internal/handle"
	"tools-back/internal/svc"

	"github.com/gin-gonic/gin"
)

// AddRouter 添加路由
func AddRouter(r *gin.Engine, svc *svc.ServerContext) *gin.Engine {
	// 设置 SSE 路由
	r.GET("/events", func(c *gin.Context) {
		handle.SSEHandler(c, svc.FrontSSEMessageChan)
	})
	v2 := r.Group("/v1/api")
	{
		v2.GET("/get/upload", func(c *gin.Context) {
			handle.OSSUploadHandler(c, svc)
		})
		v2.POST("/create/task", func(c *gin.Context) {
			handle.ToolHandler(c, svc)
		})
	}

	return r // 返回修改后的路由
}
