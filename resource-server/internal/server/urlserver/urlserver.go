package urlserver

import (
	"context"
	"resource-server/internal/svc"
)

type UrlServer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUrlServer(ctx context.Context, svcCtx *svc.ServiceContext) *UrlServer {
	return &UrlServer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
