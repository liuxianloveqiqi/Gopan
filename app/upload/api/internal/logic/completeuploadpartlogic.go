package logic

import (
	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"
	"Gopan/app/upload/rpc/types/upload"
	"Gopan/common/errorx"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type CompleteUploadPartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCompleteUploadPartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompleteUploadPartLogic {
	return &CompleteUploadPartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CompleteUploadPartLogic) CompleteUploadPart(req *types.CompleteUploadPartReq) error {
	// todo: add your logic here and delete this line
	// 获取user_id
	var err error
	userId, ok := l.ctx.Value("user_id").(int64)
	if !ok {
		return errors.Wrapf(errorx.NewDefaultError("user_id获取失败"), "user_id获取失败 user_id:%v", userId)
	}
	// 通过 uploadId 查询 Redis 并判断是否所有分块上传完成
	result, err := l.svcCtx.Rdb.HGetAll(l.ctx, "multipart_"+req.UploadID).Result()
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "redis查询错误err:%v", err)
	}
	count := 0
	// 遍历map
	for k, v := range result {
		// 检测k是否以"checkindex_"为前缀并且v为"1"
		if strings.HasPrefix(k, "checkindex_") && strconv.Itoa(1) == v {
			count++
		}
	}
	// 所需分片数量不等于redis中查出来已经完成分片的数量，返回无法满足合并条件
	if count != int(req.ChunkCount) {
		return errors.Wrapf(errorx.NewCodeError(40004, errorx.ErrMultipartUploadNoComplete), "分块上传文件的时候没有传完就调用合并分块文件接口 请求：%v", req)
	}
	// 开始合并分块
	// 合并后的文件路径
	mergedFilePath := "/Users/liuxian/GoProjects/project/Gopan/data/file/" + req.FileSha1 + "/" + req.FileName
	err = os.MkdirAll(path.Dir(mergedFilePath), 0744)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "make文件夹错误 err:%v", err)
	}

	mergedFile, err := os.Create(mergedFilePath)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError(err.Error()), "os.Create合并后的文件路径  err:%v", err)

	}
	defer mergedFile.Close()
	// 读取每个分块文件数据并加入到合并文件中
	for i := 1; i <= int(req.ChunkCount); i++ {
		chunkFilePath := "/Users/liuxian/GoProjects/project/Gopan/data/file/" + req.UploadID + "/" + strconv.Itoa(i) // 分块文件路径
		chunkData, err := os.ReadFile(chunkFilePath)
		if err != nil {
			return errors.Wrapf(errorx.NewDefaultError(err.Error()), "os.ReadFile分块文件路径 err:%v", err)

		}

		_, err = mergedFile.Write(chunkData)
		if err != nil {
			return errors.Wrapf(errorx.NewDefaultError(err.Error()), "write分片文件内容到合并文件 err:%v", err)

		}

		// 删除已合并的分块文件
		err = os.Remove(chunkFilePath)
		if err != nil {
			logc.Error(l.ctx, "无法删除已经合并的分块文件 err:", err)
		}
	}

	// 调用rpc上传合并后的文件到对应到存储
	_, err = l.svcCtx.Rpc.UploadFile(l.ctx, &upload.UploadFileReq{
		UserId:           userId,
		FileSha1:         req.FileSha1,
		FileSize:         req.FileSize,
		FileName:         req.FileName,
		FileAddr:         mergedFilePath,
		CreateTime:       timestamppb.New(time.Now()),
		UpdateTime:       timestamppb.New(time.Now()),
		Status:           0,
		CurrentStoreType: req.CurrentStoreType,
	})
	if err != nil {
		return errors.Wrapf(err, "req: %+v", req)
	}
	return nil
}
