package logic

import (
	"Gopan/app/upload/rpc/types/upload"
	"Gopan/common/errorx"
	"context"
	"github.com/pkg/errors"

	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TryFastFileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTryFastFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TryFastFileUploadLogic {
	return &TryFastFileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TryFastFileUploadLogic) TryFastFileUpload(req *types.TryFastUploadReq) error {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	_, err := l.svcCtx.Rpc.FastUploadFile(l.ctx, &upload.FastUploadFileReq{
		UserId:   userId,
		FileSha1: req.FileSha1,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}
	return nil
}
