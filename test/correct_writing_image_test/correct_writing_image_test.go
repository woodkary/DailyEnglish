package test

import (
	params "DailyEnglish/CorrectWritingRequestParams"
	"DailyEnglish/utils"
	"DailyEnglish/utils/authv4"
	"fmt"
	"testing"
)

func TestImageCorrect(t *testing.T) {
	// 添加请求参数
	paramsMap := params.GetRequestMap(params.ReadImage("C:\\Users\\karywoodOyo\\Desktop\\essays\\考研英语1小作文___graduate_a1.png"))
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv4.AddAuthParams(params.AppKey, params.AppSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/v2/correct_writing_image", header, paramsMap, "application/json")
	// 打印返回结果
	fmt.Println("result:", string(result))
	//转化为对象，并格式化输出结果
	formatResult, err := params.ParseResultFromJSON(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	params.FormatResult(formatResult)
}
