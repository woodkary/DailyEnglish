package correctWriting

import (
	"DailyEnglish/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type CorrectWritingRequestParams struct {
	Q                 []string `json:"q"`
	Grade             []string `json:"grade"`
	Title             []string `json:"title"`
	ModelContent      []string `json:"modelContent"`
	IsNeedSynonyms    []string `json:"isNeedSynonyms"`
	CorrectVersion    []string `json:"correctVersion"`
	IsNeedEssayReport []string `json:"isNeedEssayReport"`
}

func GetRequestMap(req *CorrectWritingRequestParams) map[string][]string {
	return map[string][]string{
		"q":                 req.Q,
		"grade":             req.Grade,
		"title":             req.Title,
		"modelContent":      req.ModelContent,
		"isNeedSynonyms":    req.IsNeedSynonyms,
		"correctVersion":    req.CorrectVersion,
		"isNeedEssayReport": req.IsNeedEssayReport,
	}
}

// 从文件中读取文章，并返回参数结构体
// 文件格式：fileName___grade
// 例如：test___graduate_a1.txt，即考研英语一小作文
func ReadArticle(fileName string) *CorrectWritingRequestParams {
	//通过文件名的三个下划线___之后点号.之前的内容，判断等级
	grade := fileName[strings.LastIndex(fileName, "___")+3 : strings.LastIndex(fileName, ".")]
	fmt.Println("grade:", grade)
	if grade == "" {
		grade = "default"
	}
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
	return &CorrectWritingRequestParams{
		Q:                 []string{string(content)},
		Grade:             []string{grade},
		Title:             []string{""},
		ModelContent:      []string{""},
		IsNeedSynonyms:    []string{"false"},
		CorrectVersion:    []string{"basic"},
		IsNeedEssayReport: []string{"false"},
	}
}

// 从图片读取文章
// 图片格式：fileName___grade.jpg
// 例如：test___graduate_a1.jpg，即考研英语一小作文
func ReadImage(fileName string) *CorrectWritingRequestParams {
	//如果不是图片格式，则返回空
	if !strings.HasSuffix(fileName, ".jpg") && !strings.HasSuffix(fileName, ".png") && !strings.HasSuffix(fileName, ".bmp") {
		return nil
	}
	//通过文件名的三个下划线___之后点号.之前的内容，判断等级
	grade := fileName[strings.LastIndex(fileName, "___")+3 : strings.LastIndex(fileName, ".")]
	fmt.Println("grade:", grade)
	if grade == "" {
		grade = "default"
	}
	//读取图片内容，转为base64编码
	content, err := utils.ReadFileAsBase64(fileName)
	if err != nil {
		fmt.Println("read file error:", err)
		return nil
	}
	//图片内容即为Q，其他参数全部默认值
	return &CorrectWritingRequestParams{
		Q:                 []string{string(content)},
		Grade:             []string{grade},
		Title:             []string{""},
		ModelContent:      []string{""},
		IsNeedSynonyms:    []string{"false"},
		CorrectVersion:    []string{"basic"},
		IsNeedEssayReport: []string{"false"},
	}
}
