package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func main() {
	// 腾讯云 COS 配置信息
	secretID := "your_secret_id"
	secretKey := "your_secret_key"
	bucket := "your_bucket_name"
	region := "your_region"
	baseURL := fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region)

	// 初始化 COS 客户端
	u, _ := url.Parse(baseURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	// 从命令行读取文件夹路径
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入文件夹路径: ")
	localFolderPath, _ := reader.ReadString('\n')
	localFolderPath = strings.TrimSpace(localFolderPath)

	// 上传文件夹中的所有文件到 COS
	err := filepath.Walk(localFolderPath, func(localPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过文件夹
		if info.IsDir() {
			return nil
		}

		// 读取文件内容
		fileContent, err := ioutil.ReadFile(localPath)
		if err != nil {
			return err
		}

		// 构建 COS 中的对象路径
		relativePath := strings.TrimPrefix(localPath, localFolderPath)
		cosPath := strings.ReplaceAll(relativePath, "\\", "/")

		// 上传文件到 COS
		_, err = client.Object.Put(context.Background(), cosPath, strings.NewReader(string(fileContent)), nil)
		if err != nil {
			return err
		}

		fmt.Printf("Uploaded %s to %s\n", localPath, cosPath)
		return nil
	})

	if err != nil {
		fmt.Printf("Error uploading files: %v\n", err)
	} else {
		fmt.Println("Upload complete!")
	}
}
