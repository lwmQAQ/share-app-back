package handle

import (
	"context"
	"net/http"
	"oss-server/internal/server"
	"oss-server/internal/svc"
	"oss-server/internal/types"

	"github.com/gin-gonic/gin"
)

func OSSUploadHandler(c *gin.Context, svc *svc.ServiceContext) {
	var req = new(types.OSSUploadReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	server := server.NewOssServer(context.Background(), svc)
	resp := server.OSSUpload(req)
	c.JSON(http.StatusOK, resp)
}
