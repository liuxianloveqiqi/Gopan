package test

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"testing"
)

func TestMinio(t *testing.T) {
	// 连接到 MinIO 集群
	endpoint := "43.139.195.17:7878"
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
	filePath := "/Users/liuxian/GoProjects/project/Gopan/hello.txt"

	// 创建存储桶（如果不存在）
	err = minioClient.MakeBucket(context.TODO(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.TODO(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists", bucketName)
		} else {
			log.Fatalln(err, 333)
		}
	} else {
		log.Printf("Bucket '%s' created successfully", bucketName)
	}

	// 打开本地文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// 上传文件到 MinIO
	_, err = minioClient.PutObject(context.TODO(), bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("File '%s' uploaded successfully to bucket '%s' as '%s'\n", filePath, bucketName, objectName)

}
