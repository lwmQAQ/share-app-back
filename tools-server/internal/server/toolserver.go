package server

import (
	"context"
	"tools-back/internal/models/adapter"
	"tools-back/internal/svc"
	"tools-back/internal/types"
)

type ToolServer struct {
	ctx    context.Context
	svcCtx *svc.ServerContext
}

func NewToolServer(ctx context.Context, svcCtx *svc.ServerContext) *ToolServer {
	return &ToolServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *ToolServer) Translation(req *types.TranslationReq) types.Response {

	insert := adapter.BuildInsertTask(req.DownLoadUrl, 20000, req.FileName)

	taskid, err := s.svcCtx.TaskDao.InsertTask(insert)
	if err != nil {
		s.svcCtx.Logger.Error("无法插入任务 原因为:", err)
		return types.ErrorMsg("系统错误")
	}

	//2.异步获取rpc服务
	message := &types.RpcChanMessage{
		TaskID:      taskid,
		DownLoadUrl: req.DownLoadUrl,
		FileName:    req.FileName,
	}
	*s.svcCtx.RpcChan <- message // 发送消息到通道
	return types.Success(types.TranslationResp{
		TaskID: taskid,
	})
}

func (s *ToolServer) Format() {

}
