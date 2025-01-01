package handle

import (
	"context"
	"fmt"
	"net/http"
	"tools-back/internal/server"
	"tools-back/internal/svc"
	"tools-back/internal/types"

	"github.com/gin-gonic/gin"
)

func ToolHandler(c *gin.Context, svc *svc.ServerContext) {
	var req = new(types.TranslationReq) // 初始化结构体指针

	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println(req)
	server := server.NewToolServer(context.Background(), svc)
	resp := server.Translation(req)
	c.JSON(http.StatusOK, resp)
}
