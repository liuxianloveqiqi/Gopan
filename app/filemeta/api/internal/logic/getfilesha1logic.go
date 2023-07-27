package logic

import (
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"Gopan/app/filemeta/api/internal/svc"
	"Gopan/app/filemeta/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileSha1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileSha1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileSha1Logic {
	return &GetFileSha1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileSha1Logic) GetFileSha1(r *http.Request) (resp *types.GetFileSha1Resp, err error) {
	// todo: add your logic here and delete this line
	// 接收文件流及存储到本地目录
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("Failed to get data, err:%s\n", err.Error())
		return nil, errors.Wrapf(errorx.NewCodeError(40001, errorx.ErrFileOpen), "打开文件错误 err:%v", err)
	}
	defer file.Close()

	// 计算文件sha1值
	file.Seek(0, 0)
	fileSha1 := utils.FileSha1(file)
	file.Seek(0, 0)
	return &types.GetFileSha1Resp{
		FileSha1: fileSha1,
	}, nil
}
