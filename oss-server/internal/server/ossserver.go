package server

import (
	"context"
	"oss-server/internal/svc"
	"oss-server/internal/types"
	"oss-server/utils"
	"path/filepath"
)

type OssServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssServer(ctx context.Context, svcCtx *svc.ServiceContext) *OssServer {
	return &OssServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (s *OssServer) OSSUpload(req *types.OSSUploadReq) types.Response {
	// Get the file extension (including the dot)
	ext := filepath.Ext(req.FileName)
	uploadurl, downloadurl, err := s.svcCtx.MinioClient.GetUploadUrl(s.svcCtx.ServerConfig.Minio.BucketName, utils.NewUUID(), req.Scene, ext)
	if err != nil {
		s.svcCtx.Logger.Errorf("生成上传链接失败 %v", err)
	}

	return types.Success(types.OSSUploadResp{
		UploadUrl:   uploadurl,
		DownloadUrl: downloadurl,
	})
}
