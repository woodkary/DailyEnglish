package utils

import (
	"DailyEnglish/utils/authv4"
	"encoding/json"
	"fmt"
)

// 您的应用ID
var appKey = "5a89afd315889255"

// 您的应用密钥
var appSecret = "IukrwPmugpMwRUH4Nc7AcV2LU2xxdOF1"

type MajorScore struct {
	GrammarAdvice   string  `json:"grammarAdvice"`
	WordScore       float64 `json:"wordScore"`
	GrammarScore    float64 `json:"grammarScore"`
	TopicScore      float64 `json:"topicScore"`
	Emphasis        int     `json:"emphasis"`
	WordAdvice      string  `json:"wordAdvice"`
	StructureScore  float64 `json:"structureScore"`
	StructureAdvice string  `json:"structureAdvice"`
}
type Result struct {
	RawEssay    string     `json:"rawEssay"`
	SentNum     int        `json:"sentNum"`
	StLevelCode int        `json:"stLevelCode"`
	EssayAdvice string     `json:"essayAdvice"`
	Title       string     `json:"title"`
	TotalScore  float64    `json:"totalScore"`
	MajorScore  MajorScore `json:"majorScore"`
}
type Response struct {
	RequestId string `json:"RequestId"`
	ErrorCode string `json:"errorCode"`
	Result    Result `json:"result"`
}
type CorrectWritingRequestParams struct {
	Q                 []string `json:"q"`
	Grade             []string `json:"grade"`
	Title             []string `json:"title"`
	ModelContent      []string `json:"modelContent"`
	IsNeedSynonyms    []string `json:"isNeedSynonyms"`
	CorrectVersion    []string `json:"correctVersion"`
	IsNeedEssayReport []string `json:"isNeedEssayReport"`
}

var gradeMap = map[int]string{
	1: "elementary",
	2: "junior",
	3: "high",
	4: "cet4",
	5: "cet6",
	6: "graduate",
	7: "toefl",
	8: "ielts",
	9: "academic",
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

func CallApiScore(Base64IMG string, grade int, title string) Response {
	header := map[string][]string{
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	requestParams := CorrectWritingRequestParams{
		Q:                 []string{Base64IMG},
		Grade:             []string{gradeMap[grade]},
		Title:             []string{title},
		ModelContent:      []string{""},
		IsNeedSynonyms:    []string{"false"},
		CorrectVersion:    []string{"basic"},
		IsNeedEssayReport: []string{"false"},
	}
	paramsMap := GetRequestMap(&requestParams)
	authv4.AddAuthParams(appKey, appSecret, paramsMap)
	result := DoPost("https://openapi.youdao.com/v2/correct_writing_image", header, paramsMap, "application/json")
	//把result byte数组转换为json对象
	var response Response
	err := json.Unmarshal(result, &response)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return Response{}
	}
	//返回result

	return response
}
