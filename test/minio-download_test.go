package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	objectName := "2023周二78-学习通慕课.xlsx"
	expiry := time.Duration(24) * time.Hour // 链接过期时间

	presignedURL, err := minioClient.PresignedGetObject(context.TODO(), bucketName, objectName, expiry, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("文件下载链接:", presignedURL)
}
