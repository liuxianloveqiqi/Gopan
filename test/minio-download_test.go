package test

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

func TestPame(t *testing.T) {

	endpoint := "43.139.195.17:7878"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// 创建Minio客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 获取文件的下载链接
	bucketName := "userfile"
	objectName := "docker-compose.yaml"
	filePath := "/Users/liuxian/GoProjects/project/Gopan/data/file/download/docker"
	// 使用MinIO客户端下载对象
	err = minioClient.FGetObject(context.Background(), bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("文件下载成功：", filePath)
}
