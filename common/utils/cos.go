package utils

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"time"
)

func TencentCOSUpload(urlValue, id, key, filePath string, file *os.File) error {
	//COS客户端连接
	u, _ := url.Parse(urlValue)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 10000 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  id,
			SecretKey: key,
		},
	})

	// 3.通过文件流上传对象

	_, saveErr := c.Object.Put(context.Background(), filePath, file, nil)
	if saveErr == nil {
		return nil
	} else {
		return saveErr
	}

}
func TencentCOSDownload(urlValue, id, key, bucketKey, downloadPath string) {
	//将<bucket>和<region>修改为真实的信息
	//bucket的命名规则为{name}-{appid} ，此处填写的存储桶名称必须为此格式
	//COS客户端连接
	u, _ := url.Parse(urlValue)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 10000 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  id,
			SecretKey: key,
		},
	})
	opt := &cos.MultiDownloadOptions{
		ThreadPoolSize: 5,
		CheckPoint:     true, // 开启断点续传
	}

	_, err := c.Object.Download(
		context.Background(), bucketKey, downloadPath, opt,
	)
	if err != nil {
		panic(err)
	}

}
