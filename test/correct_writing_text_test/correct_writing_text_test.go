package test

import (
	"DailyEnglish/utils"
	"DailyEnglish/utils/authv4"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// 您的应用ID
var appKey = "5a89afd315889255"

// 您的应用密钥
var appSecret = "IukrwPmugpMwRUH4Nc7AcV2LU2xxdOF1"

func TestCorrectWritingText(t *testing.T) {
	// 添加请求参数
	paramsMap := GetRequestMap(ReadArticle("C:\\Users\\karywoodOyo\\Desktop\\essays\\2022小作文___graduate_a1.txt"))
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	// 添加鉴权相关参数
	authv4.AddAuthParams(appKey, appSecret, paramsMap)
	// 请求api服务
	result := utils.DoPost("https://openapi.youdao.com/v2/correct_writing_text", header, paramsMap, "application/json")
	// 打印返回结果
	fmt.Print("result:", string(result))

}
func CreateRequestParams() map[string][]string {

	/*
		note: 将下列变量替换为需要请求的参数
		取值参考文档: https://ai.youdao.com/DOCSIRMA/html/%E4%BD%9C%E6%96%87%E6%89%B9%E6%94%B9/API%E6%96%87%E6%A1%A3/%E8%8B%B1%E8%AF%AD%E4%BD%9C%E6%96%87%E6%89%B9%E6%94%B9%EF%BC%88%E6%96%87%E6%9C%AC%E8%BE%93%E5%85%A5%EF%BC%89/%E8%8B%B1%E8%AF%AD%E4%BD%9C%E6%96%87%E6%89%B9%E6%94%B9%EF%BC%88%E6%96%87%E6%9C%AC%E8%BE%93%E5%85%A5%EF%BC%89-API%E6%96%87%E6%A1%A3.html
	*/
	q := "正文文本"
	grade := "作文等级"
	title := "作文标题"
	modelContent := "作文参考范文"
	isNeedSynonyms := "是否查询同义词"
	correctVersion := "作文批改版本：基础，高级"
	isNeedEssayReport := "是否返回写作报告"

	return map[string][]string{
		"q":                 {q},
		"grade":             {grade},
		"title":             {title},
		"modelContent":      {modelContent},
		"isNeedSynonyms":    {isNeedSynonyms},
		"correctVersion":    {correctVersion},
		"isNeedEssayReport": {isNeedEssayReport},
	}
}

type RequestParams struct {
	Q                 string `json:"q"`
	Grade             string `json:"grade"`
	Title             string `json:"title"`
	ModelContent      string `json:"modelContent"`
	IsNeedSynonyms    string `json:"isNeedSynonyms"`
	CorrectVersion    string `json:"correctVersion"`
	IsNeedEssayReport string `json:"isNeedEssayReport"`
}

func GetRequestMap(req *RequestParams) map[string][]string {
	return map[string][]string{
		"q":                 {req.Q},
		"grade":             {req.Grade},
		"title":             {req.Title},
		"modelContent":      {req.ModelContent},
		"isNeedSynonyms":    {req.IsNeedSynonyms},
		"correctVersion":    {req.CorrectVersion},
		"isNeedEssayReport": {req.IsNeedEssayReport},
	}
}

// 从文件中读取文章，并返回参数结构体
// 文件格式：fileName___grade
// 例如：test___graduate_a1.txt，即考研英语一小作文
func ReadArticle(fileName string) *RequestParams {
	//通过文件名的三个下划线___之后点号.之前的内容，判断等级
	grade := fileName[strings.LastIndex(fileName, "___")+3 : strings.LastIndex(fileName, ".")]
	fmt.Println("grade:", grade)
	//读取文件内容
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error:", err)
		return nil
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file error:", err)
		return nil
	}
	//文件内容即为Q，其他参数全部默认值
	return &RequestParams{
		Q:                 string(content),
		Grade:             grade,
		Title:             "",
		ModelContent:      "",
		IsNeedSynonyms:    "false",
		CorrectVersion:    "basic",
		IsNeedEssayReport: "false",
	}
}
