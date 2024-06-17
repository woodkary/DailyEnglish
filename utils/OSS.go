package utils

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var OssEndpoint = "************"
var OssAccessKeyId = "************"
var OssAccessKeySecret = "************"
var ossBucketName = "dailyenglish"

func GetOSSSecret(db *sql.DB) {
	// 从数据库中获取 OSS 相关信息，只获取活跃的
	rows, err := db.Query("SELECT oss_end_point, oss_access_key_id, oss_access_key_secret,is_available FROM oss_secret WHERE is_available = 1")
	if err != nil {
		fmt.Println("Failed to get OSS secret from database:", err)
		return
	}
	defer rows.Close()
	//取结果集中的第一个

	for rows.Next() {
		var ossEndPoint0, ossAccessKeyId0, ossAccessKeySecret0 string
		var isAvailable int
		err = rows.Scan(&ossEndPoint0, &ossAccessKeyId0, &ossAccessKeySecret0, &isAvailable)
		if err != nil {
			fmt.Println("Failed to scan OSS secret from database:", err)
			return
		}
		OssEndpoint = ossEndPoint0
		OssAccessKeyId = ossAccessKeyId0
		OssAccessKeySecret = ossAccessKeySecret0
		break
	}
	//解密Oss信息
	OssEndpoint = AesDecrypt(OssEndpoint, "DailyEnglish_end")
	OssAccessKeyId = AesDecrypt(OssAccessKeyId, "DailyEnglish_key")
	OssAccessKeySecret = AesDecrypt(OssAccessKeySecret, "DailyEnglish_sec")
	fmt.Println("OSS secret:", OssEndpoint, OssAccessKeyId, OssAccessKeySecret)
}

func UploadImageToOSS(base64Image string) (string, error) {
	// 创建 OSS 客户端
	client, err := oss.New(OssEndpoint, OssAccessKeyId, OssAccessKeySecret)
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
	bucket, err := client.Bucket(ossBucketName)
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
