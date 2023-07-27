package logic

import (
	"Gopan/app/filemeta/rpc/types/filemeta"
	"Gopan/common/errorx"
	"context"
	"github.com/pkg/errors"

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

	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	_, err := l.svcCtx.Rpc.UpdataUserFileMeta(l.ctx, &filemeta.UpdataUserFileMetaReq{
		UserId:   userId,
		FileName: req.FileName,
		FileSha1: req.FileSha1,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}
	return nil
}
