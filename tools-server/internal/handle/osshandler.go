package handle

import (
	"context"
	"net/http"
	"tools-back/internal/server"
	"tools-back/internal/svc"
	"tools-back/internal/types"

	"github.com/gin-gonic/gin"
)

func OSSUploadHandler(c *gin.Context, svc *svc.ServerContext) {
	var req = new(types.OssUploadReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindQuery(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	server := server.NewOssServer(context.Background(), svc)
	resp := server.Upload(req)
	c.JSON(http.StatusOK, resp)
}
