package logic

import (
	"Gopan/app/upload/model"
	"Gopan/common/batcher"
	"Gopan/common/errorx"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"time"

	"Gopan/app/upload/rpc/internal/svc"
	"Gopan/app/upload/rpc/types/upload"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	batcher *batcher.Batcher
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {

	f := &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
	// batcher配置
	options := batcher.Options{
		5,
		100,
		50,
		10 * time.Second,
	}
	// 实现batcher
	b := batcher.New(options)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.UserFile
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.UserFile))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		if err = f.svcCtx.KqPusherClient.Push(string(kd)); err != nil {
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}
	f.batcher = b
	f.batcher.Start()

	return f
	return f
}

func (l *UploadFileLogic) UploadFile(in *upload.UploadFileReq) (*upload.CommonResp, error) {
	// todo: add your logic here and delete this line
	userfile := model.UserFile{
		Id:         0,
		UserId:     in.UserId,
		FileSha1:   in.FileSha1,
		FileSize:   in.FileSize,
		FileName:   in.FileName,
		CreateTime: time.Unix(in.CreateTime.GetSeconds(), 0),
		UpdateTime: time.Unix(in.UpdateTime.GetSeconds(), 0),
		DeleteTime: gorm.DeletedAt{Time: time.Unix(in.DeleteTime.GetSeconds(), 0)},
		Status:     in.Status,
	}
	// kafka异步处理UserFile元数据,Userfile只是多个userid，所以传他到Kafka
	err := l.batcher.Add(strconv.FormatInt(in.UserId, 10), &userfile)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(40003, errorx.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}
	return &upload.CommonResp{
		Code:    0,
		Message: "success!",
	}, nil
}
