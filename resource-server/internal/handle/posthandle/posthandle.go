package posthandle

import (
	"context"
	"fmt"
	"net/http"
	"resource-server/internal/server/postserver"
	"resource-server/internal/svc"
	"resource-server/internal/types"
	"resource-server/middleware"

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
	//  参数校验
	err := middleware.ValidateStruct(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数不符合规格")
	}
	server := postserver.NewPostServer(context.Background(), svc)
	resp := server.CreatePost(req)
	c.JSON(http.StatusOK, resp)

}

func DeletePostHandle(c *gin.Context, svc *svc.ServiceContext) {

}

func LikePostHandle(c *gin.Context, svc *svc.ServiceContext) {
	postId := c.Query("postId")
	server := postserver.NewPostServer(context.Background(), svc)
	resp := server.LikePost(postId)
	c.JSON(http.StatusOK, resp)
}
