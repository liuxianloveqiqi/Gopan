package logic

import (
	"Gopan/app/upload/model"
	"Gopan/common/batcher"
	"Gopan/common/conf"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"encoding/json"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"log"
	"os"
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

func (l *UploadFileLogic) COSUpload(filePath string, file *os.File) error {
	Url := l.svcCtx.Config.TencentCOS.Url
	SecretId := l.svcCtx.Config.TencentCOS.SecretId
	SecretKey := l.svcCtx.Config.TencentCOS.SecretKey
	err := utils.TencentCOSUpload(Url, SecretId, SecretKey, filePath, file)
	if err != nil {
		logc.Error(l.ctx, "上传文件失败")
		return err
	}
	return nil
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
		100,
		1 * time.Second,
	}
	// 实现batcher
	b := batcher.New(options)
	b.Sharding = func(key string) int {
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % options.Worker
	}
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		var msgs []*model.NewUserFile
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*model.NewUserFile))
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

}

func (l *UploadFileLogic) UploadFile(in *upload.UploadFileReq) (*upload.CommonResp, error) {
	// todo: add your logic here and delete this line
	// 获取本地储存的文件
	file, err := os.Open(in.FileAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	switch in.CurrentStoreType {

	case int64(conf.StoreLocal):
		// 本地储存，不做处理
	case int64(conf.StoreMinio):
		// 文件写入Minio存储
		log.Println("开始写入minio")
		minioPath := "/minio/" + in.FileSha1 + "/" + in.FileName
		bucketName := "userfile"
		// 上传文件到 MinIO
		_, err := l.svcCtx.MinioDb.PutObject(context.TODO(), bucketName, minioPath, file, -1, minio.PutObjectOptions{})
		if err != nil {
			log.Println(err)
			return nil, errors.Wrapf(errorx.NewDefaultError("上传文件失败 err:"+err.Error()), "上传文件到minio错误 err : %v", err)
		}
		//改变存储路径
		in.FileAddr = minioPath

	case int64(conf.StoreCOS):
		// 写入COS
		log.Println("开始写入cos")
		cosPath := "/cos/" + in.FileSha1 + "/" + in.FileName
		// 写入COS
		err := l.COSUpload(cosPath, file)
		if err != nil {
			return nil, errors.Wrapf(errorx.NewDefaultError("上传文件失败 err:"+err.Error()), "上传文件到COS错误 err:%v", err)
		}
		// 改变存储路径
		in.FileAddr = cosPath
	}
	uf := model.UserFile{
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
	userfile := model.NewUserFile{UserFile: uf, FileAddr: in.FileAddr}
	// kafka异步处理file元数据
	err = l.batcher.Add(strconv.FormatInt(in.UserId, 10), &userfile)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewCodeError(40003, errorx.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}
	return &upload.CommonResp{
		Code:    0,
		Message: "success!",
	}, nil
}
