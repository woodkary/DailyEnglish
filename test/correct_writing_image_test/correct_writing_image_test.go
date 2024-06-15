package test

import (
	params "DailyEnglish/CorrectWritingRequestParams"
	"DailyEnglish/utils"
	"DailyEnglish/utils/authv4"
	"fmt"
	"testing"
)

// 您的应用ID
var appKey = "5a89afd315889255"

// 您的应用密钥
var appSecret = "IukrwPmugpMwRUH4Nc7AcV2LU2xxdOF1"

func TestImageCorrect(t *testing.T) {
	// 添加请求参数
	paramsMap := params.GetRequestMap(params.ReadImage("C:\\Users\\karywoodOyo\\Desktop\\essays\\考研英语1小作文___graduate_a1.png"))
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv4.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/v2/correct_writing_image", header, paramsMap, "application/json")
	// 打印返回结果
	fmt.Print("result:", string(result))
}
