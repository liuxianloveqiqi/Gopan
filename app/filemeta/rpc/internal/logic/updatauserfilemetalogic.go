package logic

import (
	"context"

	"Gopan/app/filemeta/rpc/internal/svc"
	"Gopan/app/filemeta/rpc/types/filemeta"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdataUserFileMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdataUserFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdataUserFileMetaLogic {
	return &UpdataUserFileMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdataUserFileMetaLogic) UpdataUserFileMeta(in *filemeta.UpdataUserFileMetaReq) (*filemeta.CommonResp, error) {
	// todo: add your logic here and delete this line

	return &filemeta.CommonResp{}, nil
}
