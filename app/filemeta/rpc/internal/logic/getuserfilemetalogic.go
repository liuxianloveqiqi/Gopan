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

type GetUserFileMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFileMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileMetaLogic {
	return &GetUserFileMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFileMetaLogic) GetUserFileMeta(in *filemeta.GetUserFileMetaReq) (*filemeta.GetUserFileMetaResp, error) {
	// todo: add your logic here and delete this line
	ms := make([]model.UserFile, 0)
	err := l.svcCtx.MysqlDb.Where("user_id = ?", in.UserId).Find(&ms).Error
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDefaultError(err.Error()), "查询userfile err:%v ", err)

	}
	userFileMetaList := make([]*filemeta.UserFileMeta, 0, len(ms))
	for _, m := range ms {
		userFileMeta := &filemeta.UserFileMeta{
			Id:         m.Id,
			FileSha1:   m.FileSha1,
			FileSize:   m.FileSize,
			FileName:   m.FileName,
			Status:     m.Status,
			CreateTime: m.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: m.UpdateTime.Format("2006-01-02 15:04:05"),
		}
		userFileMetaList = append(userFileMetaList, userFileMeta)
	}

	// Return the result
	return &filemeta.GetUserFileMetaResp{
		UserFileMetaList: userFileMetaList,
	}, nil
}
