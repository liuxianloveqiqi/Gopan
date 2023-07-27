package logic

import (
	"Gopan/app/filemeta/rpc/types/filemeta"
	"context"
	"github.com/pkg/errors"

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
	cnt, err := l.svcCtx.Rpc.GetFileMeta(l.ctx, &filemeta.GetFileMetaReq{FileSha1: req.FileSha1})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	m := types.FileMeta{
		Id:         cnt.Id,
		FileSha1:   cnt.FileSha1,
		FileSize:   cnt.FileSize,
		FileName:   cnt.FileName,
		FileAddr:   cnt.FileAddr,
		Status:     cnt.Status,
		CreateTime: cnt.CreateTime,
		UpdateTime: cnt.UpdateTime,
	}
	return &types.GetFileMetaResp{FileMeta: m}, nil
}
