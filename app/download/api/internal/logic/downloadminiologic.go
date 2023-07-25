package logic

import (
	"Gopan/common/errorx"
	"Gopan/common/utils"
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
	info, err := l.svcCtx.MinioDb.StatObject(l.ctx, bucket, object, minio.StatObjectOptions{})
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("无法获取对象信息"), "无法获取对象信息 err:%v", err)

	}

	// 计算分块的数量
	totalParts := int((info.Size + chunkSize - 1) / chunkSize)

	// 下载并合并文件
	err = l.downloadAndMergeFile(l.svcCtx.MinioDb, bucket, object, outputDir, outputFileName, totalParts)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("下载文件失败"), "下载文件失败 err:%v", err)

	}

	// 打开合并后的文件
	filePath := filepath.Join(outputDir, outputFileName)
	file, err := os.Open(filePath)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("无法打开文件"), "无法打开文件 err:%v", err)

	}
	defer file.Close()
	// 校验文件sha1
	if req.FileSha1 != utils.FileSha1(file) {
		return errors.Wrapf(errorx.NewCodeError(40004, errorx.ErrFileSha1Falsify), "err:文件sha1值校验失败文件已经被篡改:file:%v", req)
	}
	// 设置响应头
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", outputFileName))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	// 将文件内容发送给客户端
	_, err = io.Copy(w, file)
	if err != nil {
		return errors.Wrapf(errorx.NewDefaultError("无法发送文件内容"), "无法发送文件内容 err:%v", err)
	}

	// 删除已发送的合并文件
	if err := os.Remove(filePath); err != nil {
		logc.Error(l.ctx, "无法删除合并文件:", err)
	}

	return nil
}
