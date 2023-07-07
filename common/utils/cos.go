package utils

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

func TecentCOSUpload(urlvalue, id, key, filePath string, file *multipart.FileHeader) error {
	//COS客户端连接
	u, _ := url.Parse(urlvalue)
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
	fd, err := file.Open()
	if err != nil {
		return err
	}
	defer fd.Close()
	_, saveErr := c.Object.Put(context.Background(), filePath, fd, nil)
	if saveErr == nil {
		return nil
	} else {
		return saveErr
	}

}
func TecentCOSDowmload(urlvalue, id, key, filename string) {
	//将<bucket>和<region>修改为真实的信息
	//bucket的命名规则为{name}-{appid} ，此处填写的存储桶名称必须为此格式
	u, _ := url.Parse(urlvalue)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		//设置超时时间
		Timeout: 10000 * time.Second,
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  os.Getenv(id),
			SecretKey: os.Getenv(key),
		},
	})

	name := filename
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
}
