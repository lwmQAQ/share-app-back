package postserver

import (
	"context"
	"resource-server/internal/constants"
	"resource-server/internal/ecode"
	"resource-server/internal/models"
	"resource-server/internal/models/adapter"
	"resource-server/internal/rpcclient/userclient"
	"resource-server/internal/svc"
	"resource-server/internal/types"
	"resource-server/utils"
	"time"

	"github.com/sirupsen/logrus"
)

type PostServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostServer(ctx context.Context, svcCtx *svc.ServiceContext) *PostServer {
	return &PostServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *PostServer) CreatePost(req *types.CreatePostReq) types.Response {

	//1.获取rpc作者信息
	addr, err := s.svcCtx.EtcdUtil.GetServiceInstance("UserServer")
	if err != nil {
		s.svcCtx.Logger.Errorf("获取rpc服务失败 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	rpcreq := &userclient.GetUserInfoReq{
		Id: req.UserID,
	}

	resp, err := s.svcCtx.UserRpcClient.GetUserInfo(context.Background(), rpcreq, addr)
	if err != nil {
		s.svcCtx.Logger.Errorf("rpc服务出错 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	author := &models.User{
		Username: resp.Username,
		Avatar:   resp.Avatar,
		ID:       resp.Id,
	}

	//2.TODO异步增加用户经验
	//3.插入mongo文档
	post := adapter.BuildInsertPost(author, req)
	postid, err := s.svcCtx.MongoUtil.InsertDocument("Post", post)
	if err != nil {
		s.svcCtx.Logger.Errorf("mongodb出错 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	//4. 异步更新es
	go func(id string, client *utils.ESClient, logger *logrus.Logger) {
		//TODO 异步更新es
		resource := &models.Resource{
			Title:       post.Title,
			Tags:        post.Tags,
			ClickCount:  0,
			Publisher:   author.Username,
			PublishTime: time.Now(),
		}
		err := client.InsertDocument(constants.ResourceIndex, resource, id)
		if err != nil {
			logger.Errorf("es插入出错 %v", err)
		}
	}(postid.(string), s.svcCtx.ESClient, s.svcCtx.Logger)
	return types.Success(&types.CreatePostResp{})
}
