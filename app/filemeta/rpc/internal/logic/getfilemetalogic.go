package logic

import (
	"Gopan/app/filemeta/model"
	"Gopan/common/errorx"
	"context"
	"github.com/pkg/errors"

	"Gopan/app/filemeta/rpc/internal/svc"
	"Gopan/app/filemeta/rpc/types/filemeta"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileMetaLogic {
	return &GetFileMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileMetaLogic) GetFileMeta(in *filemeta.GetFileMetaReq) (*filemeta.FileMeta, error) {
	// todo: add your logic here and delete this line
	meta := model.File{}
	if err := l.svcCtx.MysqlDb.Model(&model.File{}).Where("file_Sha1 = ?", in.FileSha1).First(&meta).Error; err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "查询文件Sha1 err:%v ", err)
	}
	return &filemeta.FileMeta{
		Id:         meta.Id,
		FileSha1:   meta.FileSha1,
		FileSize:   meta.FileSize,
		FileName:   meta.FileName,
		FileAddr:   meta.FileAddr,
		Status:     meta.Status,
		CreateTime: meta.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: meta.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}
