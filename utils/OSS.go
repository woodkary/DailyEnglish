package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadImageToOSS(base64Image string) (string, error) {
	// 创建 OSS 客户端
	client, err := oss.New("***************", "*************", "*************")
	if err != nil {
		return "", err
	}

	// 获取当前时间戳
	now := time.Now()
	filename := fmt.Sprintf("%d/%02d/%02d/%d.jpg", now.Year(), now.Month(), now.Day(), now.Unix()) // 假设图片是 JPG 格式

	// 解码 base64 编码的图像
	imgBytes, err := DecodeBase64(base64Image)
	if err != nil {
		return "", err
	}

	// 创建一个新的 OSS 桶客户端
	bucket, err := client.Bucket("dailyenglish")
	if err != nil {
		return "", err
	}

	// 将图像上传到 OSS
	err = bucket.PutObject(filename, bytes.NewReader(imgBytes))
	if err != nil {
		return "", err
	}

	return filename, nil
}

func DecodeBase64(base64String string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return decodedBytes, nil
}
