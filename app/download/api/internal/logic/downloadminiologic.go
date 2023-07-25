package logic

import (
	"Gopan/common/errorx"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"Gopan/app/download/api/internal/svc"
	"Gopan/app/download/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	bucket    = "userfile"
	outputDir = "/Users/liuxian/GoProjects/project/Gopan/data/file/download/"
	chunkSize = 50 * 1024 * 1024 // 每个分块的大小（50MB）
)

type DownloadMinioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadMinioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadMinioLogic {
	return &DownloadMinioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadMinioLogic) downloadFilePart(client *minio.Client, bucket, object string, partNumber int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	// 创建文件以保存下载的分块
	filePath := outputDir + fmt.Sprintf("part%d", partNumber)

	// 设置分块下载的选项
	opts := minio.GetObjectOptions{}
	opts.PartNumber = partNumber
	// 下载分块并将其写入文件
	err := client.FGetObject(context.Background(), bucket, object, filePath, opts)
	if err != nil {
		logc.Errorf(l.ctx, "下载分块失败:", err)
		ch <- "" // 将空字符串发送到通道表示下载失败
		return
	}

	ch <- filePath // 将文件路径发送到通道表示下载成功
}

func (l *DownloadMinioLogic) mergeFileParts(outputDir, outputFileName string, totalParts int) error {
	// 创建输出文件以保存合并后的文件
	outputPath := filepath.Join(outputDir, outputFileName)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 将每个分块的内容合并到输出文件中
	for i := 0; i < totalParts; i++ {
		partPath := outputDir + fmt.Sprintf("part%d", i)
		partData, err := os.ReadFile(partPath)
		if err != nil {
			return err
		}
		_, err = outputFile.Write(partData)
		if err != nil {
			return err
		}

		// 删除已合并的分块文件
		if err := os.Remove(partPath); err != nil {
			logc.Errorf(l.ctx, "无法删除分块文件:", err)
		}
	}

	return nil
}

func (l *DownloadMinioLogic) downloadAndMergeFile(client *minio.Client, bucket, object, outputDir, outputFileName string, totalParts int) error {
	// 创建通道以接收成功下载的分块文件路径
	ch := make(chan string, totalParts)

	// 并发下载分块
	var wg sync.WaitGroup
	for partNumber := 0; partNumber < totalParts; partNumber++ {
		wg.Add(1)
		go l.downloadFilePart(client, bucket, object, partNumber, &wg, ch)
	}

	// 等待所有分块下载完成
	wg.Wait()

	// 检查分块下载是否成功
	close(ch)
	for filePath := range ch {
		if filePath == "" {
			return fmt.Errorf("下载分块失败")
		}
	}

	// 合并分块为完整文件
	err := l.mergeFileParts(outputDir, outputFileName, totalParts)
	if err != nil {
		return err
	}

	return nil
}
func (l *DownloadMinioLogic) DownloadMinio(req *types.DownloadMinioReq, w http.ResponseWriter, r *http.Request) error {
	// todo: add your logic here and delete this line
	// 获取对象信息
	fmt.Println(req)
	object := req.FileAddr
	outputFileName := req.FileName

	// Fetch the object from MinIO
	objectData, err := l.svcCtx.MinioDb.GetObject(l.ctx, bucket, object, minio.GetObjectOptions{})
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("无法获取对象信息"), "无法获取对象信息 err:%v", err)
	}
	defer objectData.Close()

	//// Create a file on the server to save the MinIO object data
	//localFilePath := "/Users/liuxian/GoProjects/project/Gopan/data/file/download/" + outputFileName
	//localFile, err := os.Create(localFilePath)
	//if err != nil {
	//	return errors.Wrapf(errorx.NewDefaultError("无法创建本地文件"), "无法创建本地文件 err:%v", err)
	//}
	//defer localFile.Close()
	//
	//// Copy the MinIO object data to the local file
	//_, err = io.Copy(localFile, objectData)
	//if err != nil {
	//	return errors.Wrapf(errorx.NewDefaultError("无法写入本地文件"), "无法写入本地文件 err:%v", err)
	//}

	// Set the response headers for streaming download

	// Calculate SHA1 hash of object content.sha1Hash = sha1.Sum(obiectContent)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", outputFileName))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	// Stream the local file data directly to the client
	_, err = io.Copy(w, objectData)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("无法发送文件内容"), "无法发送文件内容 err:%v", err)
	}

	return nil
}
