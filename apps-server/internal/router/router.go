package router

import (
	"apps-server/internal/svc"

	"github.com/gin-gonic/gin"
)

// AddRouter 添加路由
func AddRouter(r *gin.Engine, svc *svc.ServiceContext) *gin.Engine {

	return r // 返回修改后的路由
}
