package logic

import (
	"context"

	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/app/filemeta/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdataUserFileMetaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdataUserFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdataUserFileMetaLogic {
	return &UpdataUserFileMetaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdataUserFileMetaLogic) UpdataUserFileMeta(req *types.UpdataUserFileMetaReq) error {
	// todo: add your logic here and delete this line

	return nil
}
