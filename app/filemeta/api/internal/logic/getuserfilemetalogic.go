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
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	cnt, err := l.svcCtx.Rpc.GetUserFileMeta(l.ctx, &filemeta.GetUserFileMetaReq{UserId: userId})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", userId)
	}

	return &types.GetUserFileMetaResp{UserFileMetaList: cnt}, nil
}
