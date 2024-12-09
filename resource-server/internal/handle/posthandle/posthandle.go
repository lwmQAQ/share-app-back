package posthandle

import (
	"context"
	"fmt"
	"net/http"
	"resource-server/internal/server/postserver"
	"resource-server/internal/svc"
	"resource-server/internal/types"

	"github.com/gin-gonic/gin"
)

func SearchDetailsHandle(c *gin.Context, svc *svc.ServiceContext) {
	postId := c.Query("postId")
	fmt.Println(postId)
	server := postserver.NewPostServer(context.Background(), svc)
	resp := server.GetPostById(postId)
	c.JSON(http.StatusOK, resp)
}

func CreatePostHandle(c *gin.Context, svc *svc.ServiceContext) {
	var req = new(types.CreatePostReq) // 初始化结构体指针
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// TODO 参数校验
	server := postserver.NewPostServer(context.Background(), svc)
	resp := server.CreatePost(req)
	c.JSON(http.StatusOK, resp)

}

func DeletePostHandle(c *gin.Context, svc *svc.ServiceContext) {

}
