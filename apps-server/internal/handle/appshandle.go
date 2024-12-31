package handle

import (
	"apps-server/internal/server"
	"apps-server/internal/svc"
	"apps-server/internal/types"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAppsList(c *gin.Context, svc *svc.ServiceContext) {
	req := new(types.AppListReq)
	if err := c.ShouldBindQuery(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}
	server := server.NewAppServer(context.Background(), svc)
	resp := server.GetAppsList(req)
	c.JSON(http.StatusOK, resp)
}

func GetAppDetials(c *gin.Context, svc *svc.ServiceContext) {
	id := c.Param("id") // 获取路径中的参数
	// 检查id的长度是否为20
	if len(id) != 36 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID 不存在"})
		return
	}
	server := server.NewAppServer(context.Background(), svc)
	resp := server.GetAppDetials(id)
	c.JSON(http.StatusOK, resp)
}
