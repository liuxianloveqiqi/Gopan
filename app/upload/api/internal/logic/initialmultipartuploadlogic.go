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

type InitialMultipartUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitialMultipartUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitialMultipartUploadLogic {
	return &InitialMultipartUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitialMultipartUploadLogic) InitialMultipartUpload(req *types.InitialMultipartUploadReq) (resp *types.InitialMultipartUploadResp, err error) {
	// todo: add your logic here and delete this line
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	cnt, err := l.svcCtx.Rpc.InitialMultipartUpload(l.ctx, &upload.InitialMultipartUploadReq{
		UserId:   userId,
		FileSha1: req.FileSha1,
		FileSize: req.FileSize,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.InitialMultipartUploadResp{
		FileSha1:   cnt.FileSha1,
		FileSize:   cnt.FileSize,
		UploadID:   cnt.UploadID,
		ChunkSize:  cnt.ChunkSize,
		ChunkCount: cnt.ChunkCount,
		UserId:     cnt.UserId,
	}, nil
}
