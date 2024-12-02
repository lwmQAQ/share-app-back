package userhandler

import (
	"context"
	"login-server/internal/ecode"
	"login-server/internal/server/userserver"
	"login-server/internal/svc"
	"login-server/internal/types"
	"login-server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfoHandler(c *gin.Context, svc *svc.ServiceContext) {
	userId, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, types.Error(ecode.ErrUserNotExist))
		return
	}

	userIdUint64, ok := userId.(uint64)
	if !ok {
		c.JSON(http.StatusInternalServerError, types.Error(ecode.ErrUserNotExist))
		return
	}

	server := userserver.NewUserServer(context.Background(), svc)
	resp := server.GetUserinfo(userIdUint64)
	c.JSON(http.StatusOK, resp)
}

func LoginByCoderHandler(c *gin.Context, svc *svc.ServiceContext) {
	ip := c.ClientIP()
	var req = new(types.LoginCodeReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//参数校验
	if err := middleware.ValidateStruct(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}
	server := userserver.NewUserServer(context.Background(), svc)
	resp := server.LoginByCode(req, ip)

	c.JSON(http.StatusOK, resp)
}

func LoginHandler(c *gin.Context, svc *svc.ServiceContext) {
	ip := c.ClientIP()
	var req = new(types.LoginReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//参数校验
	if err := middleware.ValidateStruct(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}
	server := userserver.NewUserServer(context.Background(), svc)
	resp := server.Login(req, ip)

	c.JSON(http.StatusOK, resp)
}

func RegisterHandler(c *gin.Context, svc *svc.ServiceContext) {
	ip := c.ClientIP()
	var req = new(types.RegisterReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//参数校验
	if err := middleware.ValidateStruct(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}
	server := userserver.NewUserServer(context.Background(), svc)
	resp := server.Register(req, ip)
	c.JSON(http.StatusOK, resp)
}

func SendCodeHandler(c *gin.Context, svc *svc.ServiceContext) {
	var req = new(types.SendCodeReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindQuery(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//参数校验
	if err := middleware.ValidateStruct(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}
	server := userserver.NewUserServer(context.Background(), svc)
	resp := server.SendCode(req.Email)
	c.JSON(http.StatusOK, resp)
}

func UpdateUserHandler(c *gin.Context, svc *svc.ServiceContext) {
	var req = new(types.UpdateUserReq) // 初始化结构体指针
	// 1. 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//参数校验
	if err := middleware.ValidateStruct(req); err != nil {
		svc.Logger.Errorf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数不合法"})
		return
	}
	server := userserver.NewUserServer(context.Background(), svc)
	resp := server.UpdateUser(req)
	c.JSON(http.StatusOK, resp)
}
