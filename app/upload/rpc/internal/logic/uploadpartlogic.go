package logic

import (
	"context"

	"Gopan/app/upload/rpc/internal/svc"
	"Gopan/app/upload/rpc/types/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartLogic {
	return &UploadPartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadPartLogic) UploadPart(in *upload.UploadFileReq) (*upload.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &upload.CommonResp{}, nil
}
