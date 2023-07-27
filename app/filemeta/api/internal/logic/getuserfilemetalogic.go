package logic

import (
	"context"

	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/app/filemeta/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileMetaLogic {
	return &GetUserFileMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFileMetaLogic) GetUserFileMeta() (resp *types.GetUserFileMetaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
