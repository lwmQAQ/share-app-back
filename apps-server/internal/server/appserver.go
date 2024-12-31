package server

import (
	"apps-server/internal/svc"
	"apps-server/internal/types"
	"context"
)

type AppServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppServer(ctx context.Context, svcCtx *svc.ServiceContext) *AppServer {
	return &AppServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *AppServer) GetAppsList(req *types.AppListReq) types.Response {
	// 查询数据库
	resp, err := s.svcCtx.AppsDao.GetAppsList(req.Type)
	if err != nil {
		s.svcCtx.Logger.Error("查询数据库失败 原因:", err)
		return types.ErrorMsg("系统错误") // 假设有一个通用的错误处理方法
	}

	// 初始化应用列表
	var applist []types.AppList

	// 将数据库查询结果转换为响应结构体
	for _, app := range *resp {
		temp := types.AppList{
			Name:        app.Name,
			ID:          app.ID,
			Description: app.Description,
			Icon:        app.Icon,
			Url:         app.Url,
		}
		applist = append(applist, temp) // 添加到列表中
	}

	// 返回成功响应
	return types.Success(&types.AppListResp{
		AppsList: &applist,
	})
}

func (s *AppServer) GetAppDetials(id string) types.Response {
	resp, err := s.svcCtx.AppsDao.GetAppDetials(id)
	if err != nil {
		s.svcCtx.Logger.Error("查询数据库失败 原因:", err)
		return types.ErrorMsg("系统错误") // 假设有一个通用的错误处理方法
	}
	return types.Success(resp)
}
