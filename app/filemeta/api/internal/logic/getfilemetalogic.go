package logic

import (
	"context"

	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/app/filemeta/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileMetaLogic {
	return &GetFileMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileMetaLogic) GetFileMeta(req *types.GetFileMetaReq) (resp *types.GetFileMetaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
