package logic

import (
	"Gopan/app/upload/api/internal/svc"
	"Gopan/app/upload/api/internal/types"
	"Gopan/app/upload/model"
	"Gopan/common/batcher"
	"Gopan/common/conf"
	"Gopan/common/errorx"
	"Gopan/common/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type FileUploadLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	batcher *batcher.Batcher
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	f := &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
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
}

func (l *FileUploadLogic) COSUpload(filePath string, file *multipart.FileHeader) error {
	Url := l.svcCtx.Config.TecentCOS.Url
	SecretId := l.svcCtx.Config.TecentCOS.SecretId
	SecretKey := l.svcCtx.Config.TecentCOS.SecretKey
	err := utils.TecentCOSUpload(Url, SecretId, SecretKey, filePath, file)
	if err != nil {
		logc.Error(l.ctx, "上传文件失败")
		return err
	}
	return nil
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line
	// 获取user_id
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
	// 添加文件元数据
	//fileMeta := model.FileMeta{
	//	FileName: head.Filename,
	//	Location: "/Users/liuxian/GoProjects/project/Gopan/tmp/" + head.Filename,
	//	UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	//}
	fileMeta := model.File{
		FileName:   head.Filename,
		FileSize:   head.Size,
		FileAddr:   "/Users/liuxian/GoProjects/project/Gopan/tmp/" + head.Filename,
		Status:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

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

	newFile.Seek(0, 0)
	// 计算文件sha1值
	fileMeta.FileSha1 = utils.FileSha1(newFile)
	// 偏移量重置
	newFile.Seek(0, 0)
	switch req.CurrentStoreType {
	case int64(conf.StoreLocal):
		// 文件写入Minio存储
	case int64(conf.StoreMinio):

		//data, _ := io.ReadAll(newFile)
		log.Println("开始写入minio")
		minioPath := "minio/" + fileMeta.FileSha1
		bucketName := "userfile"
		// 上传文件到 MinIO
		_, err = l.svcCtx.MinioDb.PutObject(context.TODO(), bucketName, fileMeta.FileName, newFile, -1, minio.PutObjectOptions{})
		if err != nil {
			log.Println(err)
			return errors.Wrapf(errorx.NewDefaultError("上传文件失败 err:"+err.Error()), "上传文件到minio错误 err : %v", err)
		}
		fileMeta.FileAddr = minioPath

	case int64(conf.StoreCOS):
		// 写入COS
		log.Println("开始写入cos")
		cosPath := "cos/" + fileMeta.FileSha1
		// 写入COS
		err = l.COSUpload(cosPath, head)
		if err != nil {
			return errors.Wrapf(errorx.NewDefaultError("上传文件失败 err:"+err.Error()), "上传文件到COS错误 err:%v", err)
		}
		fileMeta.FileAddr = cosPath

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

	// kafka异步处理UserFile元数据,Userfile只是多个userid，所以传他到Kafka
	err = l.batcher.Add(strconv.FormatInt(userId, 10), &userfile)
	if err != nil {
		return errors.Wrapf(errorx.NewCodeError(40003, errorx.ErrKafkaUserFileMeta+err.Error()), "kafka异步UserFileMeta失败 err:%v", err)
	}
	//if err = l.svcCtx.MysqlDb.Create(&fileMeta).Error; err != nil {
	//	return errors.Wrapf(errorx.NewCodeError(20002, errorx.ErrFileCreat+err.Error()), "file表Creat错误 err:%v\", err")
	//}
	//
	//if err = l.svcCtx.MysqlDb.Create(&userfile).Error; err != nil {
	//	return errors.Wrapf(errorx.NewCodeError(20003, errorx.ErrFileCreat+err.Error()), "userfile表Creat错误 err:%v\", err")
	//}
	return nil
}
