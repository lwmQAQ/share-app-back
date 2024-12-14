package commenthandle

import (
	"context"
	"net/http"
	"resource-server/internal/server/commentserver"
	"resource-server/internal/svc"
	"resource-server/internal/types"

	"github.com/gin-gonic/gin"
)

func CreateCommentHandle(c *gin.Context, svc *svc.ServiceContext) {
	var req = new(types.CreateCommentReq) // 初始化结构体指针
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// TODO 参数校验
	server := commentserver.NewCommentServer(context.Background(), svc)
	resp := server.CreateComment(req, req.UserID)
	c.JSON(http.StatusOK, resp)
}
