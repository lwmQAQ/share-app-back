package v1

import (
	"fmt"
	"resource-server/internal/router"
	"resource-server/internal/svc"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServerStart(svc *svc.ServiceContext) {
	r := gin.Default()
	// 使用 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},                                       // 允许的源
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders: []string{"X-My-Custom-Header"},                      // 可跨域访问的响应头
	}))
	router.AddRouter(r, svc)
	addr := fmt.Sprintf("%s:%d", svc.ServerConfig.Server.Host, svc.ServerConfig.Server.Port)
	r.Run(addr) // 运行服务器
}
