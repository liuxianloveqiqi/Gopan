package logic

import (
	"Gopan/app/upload/model"
	"Gopan/common/errorx"
	"context"
	"github.com/pkg/errors"

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
	// 获取user_id

	if err := l.svcCtx.MysqlDb.Model(&model.UserFile{}).Where("user_id = ? and file_sha1 = ?", in.UserId, in.FileSha1).Updates(&model.UserFile{FileName: in.FileName}).Error; err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "更新文件FileName err:%v ", err)
	}
	return &filemeta.CommonResp{}, nil
}
