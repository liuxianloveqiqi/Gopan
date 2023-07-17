package logic

import (
	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"
	"Gopan/app/upload/model"
	"Gopan/app/upload/rpc/types/upload"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"

	"net/http"

	"time"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	f := &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
	return f
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line
	// 获取user_id
	var err error
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	// 接收文件流及存储到本地目录
	file, head, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("Failed to get data, err:%s\n", err.Error())
		return errors.Wrapf(errorx.NewCodeError(40001, errorx.ErrFileOpen), "打开文件错误 err:%v", err)
	}
	defer file.Close()

	// 计算文件sha1值
	file.Seek(0, 0)
	fileSha1 := utils.FileSha1(file)
	file.Seek(0, 0)

	err = os.MkdirAll(l.svcCtx.Config.FileLocalPath+fileSha1, 0755)
	if err != nil {
		log.Fatal(err)
	}
	fileMeta := model.File{
		FileName: head.Filename,
		FileSize: head.Size,
		// 存储路径，sha1+name
		FileAddr:   l.svcCtx.Config.FileLocalPath + fileSha1 + "/" + head.Filename,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	fileMeta.FileSha1 = fileSha1
	newFile, err := os.Create(fileMeta.FileAddr)
	if err != nil {
		fmt.Printf("Failed to create file, err:%s\n", err.Error())
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "创建文件句柄错误 err:%v", err)

	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "io.copy 文件失败 err:%v", err)
	}

	// 使用kafka对Mysql存储的文件元信息进行异步处理
	userfile := model.UserFile{
		UserId:     userId,
		FileSha1:   fileMeta.FileSha1,
		FileSize:   fileMeta.FileSize,
		FileName:   fileMeta.FileName,
		Status:     fileMeta.Status,
		CreateTime: fileMeta.CreateTime,
		UpdateTime: fileMeta.UpdateTime,
	}
	// 调用rpc
	_, err = l.svcCtx.Rpc.UploadFile(l.ctx, &upload.UploadFileReq{
		UserId:           userfile.UserId,
		FileSha1:         userfile.FileSha1,
		FileSize:         userfile.FileSize,
		FileName:         userfile.FileName,
		FileAddr:         fileMeta.FileAddr,
		CreateTime:       timestamppb.New(userfile.CreateTime),
		UpdateTime:       timestamppb.New(userfile.UpdateTime),
		DeleteTime:       timestamppb.New(userfile.DeleteTime.Time),
		Status:           userfile.Status,
		CurrentStoreType: req.CurrentStoreType,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}

	return nil
}
