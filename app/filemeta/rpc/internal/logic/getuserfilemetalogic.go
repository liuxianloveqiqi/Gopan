package logic

import (
	"context"

	"Gopan/app/filemeta/rpc/internal/svc"
	"Gopan/app/filemeta/rpc/types/filemeta"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileMetaLogic {
	return &GetUserFileMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFileMetaLogic) GetUserFileMeta(in *filemeta.GetUserFileMetaReq) (*filemeta.GetUserFileMetaResp, error) {
	// todo: add your logic here and delete this line

	return &filemeta.GetUserFileMetaResp{}, nil
}
