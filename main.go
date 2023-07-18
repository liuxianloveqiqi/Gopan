package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	endpoint        = "43.139.195.17:7878"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	useSSL          = false
	bucketName      = "userfile"
	objectName      = "《思想道德与法治(2021版)》电子教材.docx"
	chunkSize       = 5 * 1024 * 1024                                               // 分块大小（5MB）
	outputDir       = "/Users/liuxian/GoProjects/project/Gopan/data/file/download/" // 合并后文件保存的目录
)

func downloadFilePart(client *minio.Client, bucket, object string, partNumber int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	// 创建文件以保存下载的分块
	filePath := fmt.Sprintf("/Users/liuxian/GoProjects/project/Gopan/data/file/download/part%d", partNumber)

	// 设置分块下载的选项
	opts := minio.GetObjectOptions{}
	opts.PartNumber = partNumber

	// 下载分块并将其写入文件
	err := client.FGetObject(context.Background(), bucket, object, filePath, opts)
	if err != nil {
		log.Printf("下载分块%d失败: %v\n", partNumber, err)
		ch <- "" // 将空字符串发送到通道表示下载失败
		return
	}

	ch <- filePath // 将文件路径发送到通道表示下载成功
}

func mergeFileParts(outputDir, outputFileName string, totalParts int) error {
	// 创建输出文件以保存合并后的文件
	outputPath := filepath.Join(outputDir, outputFileName)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// 将每个分块的内容合并到输出文件中
	for i := 0; i < totalParts; i++ {
		partPath := fmt.Sprintf("/Users/liuxian/GoProjects/project/Gopan/data/file/download/part%d", i)
		partData, err := os.ReadFile(partPath)
		if err != nil {
			return err
		}
		_, err = outputFile.Write(partData)
		if err != nil {
			return err
		}

		// 删除已合并的分块文件
		//if err := os.Remove(partPath); err != nil {
		//	log.Println("无法删除分块文件:", err)
		//}
	}

	return nil
}

func main() {
	// 初始化MinIO客户端对象
	ctx := context.Background()
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("无法创建MinIO客户端:", err)
	}

	// 获取对象信息
	info, err := minioClient.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		log.Fatalln("无法获取对象信息:", err)
	}

	// 计算分块的数量
	totalParts := int((info.Size + chunkSize - 1) / chunkSize)
	fmt.Println("分块数量：", totalParts)
	// 创建通道以接收成功下载的分块文件路径
	ch := make(chan string, totalParts)

	// 并发下载分块
	var wg sync.WaitGroup
	for partNumber := 0; partNumber < totalParts; partNumber++ {
		wg.Add(1)
		go downloadFilePart(minioClient, bucketName, objectName, partNumber, &wg, ch)
	}

	// 等待所有分块下载完成
	wg.Wait()

	// 检查分块下载是否成功
	close(ch)
	for filePath := range ch {
		if filePath == "" {
			log.Fatalln("分块下载失败")
		}
	}

	// 合并分块为完整文件
	err = mergeFileParts(outputDir, objectName, totalParts)
	if err != nil {
		log.Fatalln("合并文件失败:", err)
	}

	fmt.Println("文件下载并合并成功:", filepath.Join(outputDir, objectName))
}
