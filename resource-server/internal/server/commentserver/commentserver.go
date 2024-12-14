package commentserver

import (
	"context"
	"fmt"
	"resource-server/internal/ecode"
	"resource-server/internal/models"
	"resource-server/internal/models/adapter"
	"resource-server/internal/rpcclient/userclient"
	"resource-server/internal/svc"
	"resource-server/internal/types"
)

type CommentServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentServer(ctx context.Context, svcCtx *svc.ServiceContext) *CommentServer {
	return &CommentServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *CommentServer) GetCommentById(commendId uint64) (*models.Comment, error) {
	var comment = new(models.Comment)

	err := s.svcCtx.MongoUtil.SearchDocumentByID("Comment", commendId, comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

// 创建评论
func (s *CommentServer) CreateComment(req *types.CreateCommentReq, userId uint64) types.Response {
	var level = 1
	var path = string(req.PostID)
	if req.ParentID != nil {
		parrentcomment, err := s.GetCommentById(*req.ParentID)
		if err != nil {
			return types.Error(ecode.ErrSystemError)
		}
		level = parrentcomment.Level + 1
		path = fmt.Sprintf("%s.%d", parrentcomment.Path, parrentcomment.ID)

	}

	addr, err := s.svcCtx.EtcdUtil.GetServiceInstance("UserServer")
	if err != nil {
		s.svcCtx.Logger.Errorf("获取rpc服务失败 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	rpcreq := &userclient.GetUserInfoReq{
		Id: userId,
	}

	resp, err := s.svcCtx.UserRpcClient.GetUserInfo(context.Background(), rpcreq, addr)
	if err != nil {
		s.svcCtx.Logger.Errorf("rpc服务出错 %v", err)
		return types.Error(ecode.ErrSystemError)
	}
	user := &models.User{
		Username: resp.Username,
		Avatar:   resp.Avatar,
		ID:       resp.Id,
	}
	_, err = s.svcCtx.MongoUtil.InsertDocument("Comment", adapter.BuildInsertComment(*req, user, level, path))
	if err != nil {
		return types.Error(ecode.ErrSystemError)
	}
	return types.Success(&types.CreateCommentResp{})
}
