package logic

import (
	"Gopan/common/errorx"
	"context"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"path"
	"strconv"

	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadPartReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPartReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartReqLogic {
	return &UploadPartReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPartReqLogic) UploadPartReq(req *types.UploadPartReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line
	// 获得文件句柄，用于存储分块内容
	filepath := "/data/" + req.UploadID + "/" + strconv.FormatInt(req.ChunkIndex, 10)
	err := os.MkdirAll(path.Dir(filepath), 0744)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "make文件夹错误 err:%v", err)

	}

	fd, err := os.Create(filepath)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "creat文件错误 err:%v", err)

	}
	// 创建一个1MB大小的缓冲区
	buf := make([]byte, 1*1024*1024)
	for {
		// 从请求的Body中读取数据到缓冲区
		n, err := r.Body.Read(buf)
		// 将缓冲区中的数据写入分块文件
		fd.Write(buf[:n])
		// 读完退出循环
		if err != nil {
			break
		}
	}
	// 调用rpc更新redis关于分块文件状态
	return nil
}
