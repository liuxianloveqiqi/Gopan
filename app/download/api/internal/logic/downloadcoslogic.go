package logic

import (
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"io"
	"net/http"
	"os"
	"path"

	"Gopan/app/download/api/internal/svc"
	"Gopan/app/download/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCOSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadCOSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCOSLogic {
	return &DownloadCOSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadCOSLogic) DownloadCOS(req *types.DownloadCOSReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line
	// 缓存本机的文件位置
	localFilePath := l.svcCtx.Config.FileLocalPath + "download" + "/" + req.FileName
	// 先创建缓存本机的文件夹
	err := os.MkdirAll(path.Dir(localFilePath), 0744)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "make文件夹错误 err:%v", err)
	}
	// 调用COS进行下载
	utils.TencentCOSDownload(l.svcCtx.Config.TencentCOS.Url, l.svcCtx.Config.TencentCOS.SecretId, l.svcCtx.Config.TencentCOS.SecretKey, req.FileAddr, localFilePath)

	file, err := os.Open(localFilePath)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "Open file err:%v", err)

	}
	defer file.Close()
	// 根据sha1值进行文件校验
	if req.FileSha1 != utils.FileSha1(file) {
		return errors.Wrapf(errorx.NewCodeError(40004, errorx.ErrFileSha1Falsify), "err:文件sha1值校验失败文件已经被篡改:file:%v", req)
	}
	// 设置响应头
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", req.FileName))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	// 将文件内容发送给客户端
	_, err = io.Copy(w, file)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("无法发送文件内容"), "无法发送文件内容,err:%v", err)
	}
	// 删除已发送的合并文件
	if err := os.Remove(localFilePath); err != nil {
		logc.Error(l.ctx, "无法删除合并文件:", err)

	}
	return nil
}
