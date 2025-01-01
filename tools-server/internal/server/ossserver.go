package server

import (
	"context"
	"fmt"
	"tools-back/internal/constant"
	"tools-back/internal/svc"
	"tools-back/internal/types"

	"github.com/google/uuid"
)

type OssServer struct {
	ctx    context.Context
	svcCtx *svc.ServerContext
}

func NewOssServer(ctx context.Context, svcCtx *svc.ServerContext) *OssServer {
	return &OssServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (s *OssServer) Upload(req *types.OssUploadReq) types.Response {
	fmt.Println(req.Postfix)
	resp, err := s.svcCtx.MinioClient.GetUploadUrl(req.Postfix, constant.TOOLBUCKETNAME, CreateUUID())
	if err != nil {
		s.svcCtx.Logger.Error("无法连接Minio服务 原因:", err)
		return types.ErrorMsg("系统错误")
	}
	return types.Success(resp)
}

func CreateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}
