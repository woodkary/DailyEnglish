package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func TestOSSConnection(t *testing.T) {
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "LTAI5tLK4sabvdPKd3tkvhLw", "sOMarRA50S8Dw6UCs4dC73XkC3CNNd")
	if err != nil {
		t.Errorf("Error creating OSS client: %v", err)
		return
	}

	bucket, err := client.Bucket("dailyenglish")
	if err != nil {
		t.Errorf("Error getting bucket: %v", err)
		return
	}

	time := time.Now()
	objectName := fmt.Sprintf("images/%d/%02d/%02d/%d.jpg", time.Year(), time.Month(), time.Day(), time.Unix())
	t.Logf("objectName: %s", objectName)
	localFileName := "D:/code/DailyEnglish/static/image/登录背景.jpg"
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		t.Errorf("Error uploading file: %v", err)
		return
	}

	err = bucket.SetObjectACL(objectName, oss.ACLDefault)
	if err != nil {
		t.Errorf("Error setting ACL: %v", err)
		return
	}

	t.Logf("Connected to OSS successfully")
}
