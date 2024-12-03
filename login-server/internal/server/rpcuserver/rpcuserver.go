package rpcuserver

import (
	"context"
	"fmt"
	rpc "login-server/internal/rpcclient/userserver"
	"login-server/internal/svc"
)

// 实现 Arith 服务
type RpcUserServer struct {
	svcCtx *svc.ServiceContext
	rpc.UserServiceServer
}

func NewRpcUserServer(svc *svc.ServiceContext) *RpcUserServer {
	return &RpcUserServer{
		svcCtx: svc,
	}
}

func (s *RpcUserServer) GetUserInfo(ctx context.Context, req *rpc.GetUserInfoReq) (*rpc.GetUserInfoResp, error) {
	user, err := s.svcCtx.UserInfoCache.Get(uint64(req.Id))
	if err != nil {
		s.svcCtx.Logger.Errorln("查询用户失败", err)
		return nil, fmt.Errorf("用户不存在")
	}
	return &rpc.GetUserInfoResp{
		Username: user.Name,
		Avatar:   user.Avatar,
		Sex:      int32(user.Sex),
		Id:       user.ID,
	}, nil
}

func (s *RpcUserServer) BatchGetUserInfo(ctx context.Context, req *rpc.BatchGetUserInfoReq) (*rpc.BatchGetUserInfoResp, error) {
	users, err := s.svcCtx.UserInfoCache.GetBatch(req.Ids)
	if err != nil {
		s.svcCtx.Logger.Errorln("查询用户失败", err)
		return nil, fmt.Errorf("用户不存在")
	}
	var userinfos []*rpc.GetUserInfoResp
	for _, user := range users {
		tem := &rpc.GetUserInfoResp{
			Username: user.Name,
			Avatar:   user.Avatar,
			Sex:      int32(user.Sex),
			Id:       user.ID,
		}
		userinfos = append(userinfos, tem)
	}
	return &rpc.BatchGetUserInfoResp{
		Users: userinfos,
	}, nil

}
