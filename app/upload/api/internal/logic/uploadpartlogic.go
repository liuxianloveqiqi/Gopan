package logic

import (
	"Gopan/app/upload/rpc/types/upload"
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

type UploadPartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadPartLogic {
	return &UploadPartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadPartLogic) UploadPart(req *types.UploadPartReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line

	// 获得文件句柄，用于存储分块内容
	filepath := l.svcCtx.Config.FileLocalPath + req.UploadID + "/" + strconv.FormatInt(req.ChunkIndex, 10)
	err := os.MkdirAll(path.Dir(filepath), 0744)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "make文件夹错误 err:%v", err)
	}

	fd, err := os.Create(filepath)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "creat文件错误 err:%v", err)

	}
	defer fd.Close()
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
	_, err = l.svcCtx.Rpc.UploadPart(l.ctx, &upload.UploadPartReq{
		UploadID:   req.UploadID,
		ChunkIndex: req.ChunkIndex,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}
	return nil

}
