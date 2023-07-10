package logic

import (
	"context"

	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPartReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPartReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartReqLogic {
	return &UploadPartReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPartReqLogic) UploadPartReq(req *types.UploadPartReq) error {
	// todo: add your logic here and delete this line

	return nil
}
