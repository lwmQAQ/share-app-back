package v1

import (
	"context"
	"fmt"
	"tools-back/internal/enum"
	"tools-back/internal/models"
	"tools-back/internal/models/adapter"
	"tools-back/internal/router"
	"tools-back/internal/rpcclient/toolservice"
	"tools-back/internal/svc"
	"tools-back/internal/types"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	// 使用 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},                                       // 允许的源
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders: []string{"X-My-Custom-Header"},                      // 可跨域访问的响应头
	}))
	rpcChan := make(chan *types.RpcChanMessage, 10)
	svc := svc.NewServerContext(&rpcChan)
	RpcMessageHandle(svc)
	router.AddRouter(r, svc)
	addr := fmt.Sprintf("%s:%d", svc.ServerConfig.Server.Host, svc.ServerConfig.Server.Port)
	r.Run(addr) // 运行服务器
}

func RpcMessageHandle(ServerContext *svc.ServerContext) {
	go func(svc *svc.ServerContext) {
		for message := range *svc.RpcChan { // 从通道中接收消息
			fmt.Println(message)
			addr, err := svc.EtcdUtil.GetServiceInstance("ToolServer")
			if err != nil {
				svc.Logger.Errorf("获取rpc服务失败: %v", err)
				//更新数据库
				newTask := &models.Task{
					Status:    int(enum.Error),
					ID:        uint(message.TaskID),
					FileName:  message.FileName,
					SourceURL: message.DownLoadUrl,
				}
				updates := adapter.BuildUpdateTask(newTask)
				err := svc.TaskDao.UpdateTask(updates, uint(message.TaskID))
				if err != nil {
					svc.Logger.Error("更新数据库失败", err)
					SendErrorSSEChanMsg(ServerContext.FrontSSEMessageChan, newTask)
					continue
				}
				SendErrorSSEChanMsg(ServerContext.FrontSSEMessageChan, newTask)
				continue // 跳过本次消息，继续处理后续消息
			}
			rpcreq := &toolservice.TranslationRequest{
				DownloadUrl: message.DownLoadUrl,
			}
			resp, err := svc.ToolRpcClient.Translation(context.Background(), rpcreq, addr)
			if err != nil {
				svc.Logger.Errorf("rpc服务出错 %v", err)
				//更新数据库
				newTask := &models.Task{
					Status:    int(enum.Error),
					ID:        uint(message.TaskID),
					FileName:  message.FileName,
					SourceURL: message.DownLoadUrl,
				}
				updates := adapter.BuildUpdateTask(newTask)
				err := svc.TaskDao.UpdateTask(updates, uint(message.TaskID))
				if err != nil {
					svc.Logger.Error("更新数据库失败", err)
					SendErrorSSEChanMsg(ServerContext.FrontSSEMessageChan, newTask)
					continue
				}
				SendErrorSSEChanMsg(ServerContext.FrontSSEMessageChan, newTask)
				continue // 跳过本次消息，继续处理后续消息
			}
			// 通知前端
			newTask := &models.Task{
				MonoURL:   resp.MonoUrl,
				DualURL:   resp.DualUrl,
				Status:    int(enum.Complete),
				ID:        uint(message.TaskID),
				SourceURL: message.DownLoadUrl,
				FileName:  message.FileName,
			}
			updates := adapter.BuildUpdateTask(newTask)
			err = svc.TaskDao.UpdateTask(updates, uint(message.TaskID))
			if err != nil {
				svc.Logger.Error("更新数据库失败", err)
				SendErrorSSEChanMsg(ServerContext.FrontSSEMessageChan, newTask)
				continue
			}
			SendSucceedSSEChanMsg(ServerContext.FrontSSEMessageChan, newTask)
		}
		svc.Logger.Infof("RpcChan 已关闭，停止消息处理")
	}(ServerContext)
}

func SendErrorSSEChanMsg(FrontSSEMessageChan *chan *types.TranslationTaskResp, newTask *models.Task) {
	msg := &types.TranslationTaskResp{
		Status:    int(enum.Error),
		TaskID:    int(newTask.ID),
		SourceURL: newTask.SourceURL,
		MonoURL:   newTask.MonoURL,
		DualURL:   newTask.DualURL,
		FileName:  newTask.FileName,
	}
	*FrontSSEMessageChan <- msg
}

func SendSucceedSSEChanMsg(FrontSSEMessageChan *chan *types.TranslationTaskResp, newTask *models.Task) {
	msg := &types.TranslationTaskResp{
		Status:    int(enum.Complete),
		TaskID:    int(newTask.ID),
		SourceURL: newTask.SourceURL,
		MonoURL:   newTask.MonoURL,
		DualURL:   newTask.DualURL,
		FileName:  newTask.FileName,
	}
	*FrontSSEMessageChan <- msg
}
