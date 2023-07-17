package logic

import (
	"context"
	"net/http"

	"Gopan/app/download/api/internal/svc"
	"Gopan/app/download/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadMinioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadMinioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadMinioLogic {
	return &DownloadMinioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadMinioLogic) DownloadMinio(req *types.DownloadMinioReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line

	return nil
}
