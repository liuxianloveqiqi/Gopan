package test

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

func TestDown(t *testing.T) {
	// 连接到 MinIO 集群
	endpoint := "localhost:9009"
	accessKey := "minioadmin"
	secretKey := "minioadmin"
	useSSL := false

	// 创建一个MinIO客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// 设置存储桶名称和文件名称
	bucketName := "hehe"
	objectName := "hello.txt"
	filePath := "/Users/liuxian/GoProjects/project/Gopan/1-download.txt"
	// 下载文件
	err = minioClient.FGetObject(context.TODO(), bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("File '%s' downloaded successfully from bucket '%s' as '%s'\n", objectName, bucketName, filePath)
}
