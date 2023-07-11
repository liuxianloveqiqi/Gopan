package logic

import (
	"context"
	"fmt"

	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	// 通过redis查看是否所以分片上传完成
	// 通过 uploadid 查询 Redis 并判断是否所有分块上传完成
	upid := "your-upload-id"
	result, err := l.svcCtx.Rdb.HGetAll(l.ctx, "MP_"+upid).Result()
	if err != nil {
		fmt.Println("Complete upload failed:", err)

	}

	// 处理查询结果
	if len(result) == 0 {
		fmt.Println("Upload not found or not complete")
	} else {
		fmt.Println("Upload completed:", result)
	}
	return nil
}
