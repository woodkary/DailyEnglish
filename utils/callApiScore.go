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

type Response struct {
	RequestId string `json:"RequestId"`
	ErrorCode string `json:"errorCode"`
	Result    Result `json:"result"`
}
type Result struct {
	RawEssay      string        `json:"rawEssay"`
	EssayAdvice   string        `json:"essayAdvice"`
	Title         string        `json:"title"`
	TotalScore    float64       `json:"totalScore"`
	EssayFeedback EssayFeedback `json:"essayFeedback"`
	MajorScore    MajorScore    `json:"majorScore"`
}

type EssayFeedback struct {
	SentsFeedback []SentFeedback `json:"sentsFeedback"`
}

/*
{"grammarAdvice": "熟练使用各种语法结构写作，偶尔有语法错误，超厉害！",
"wordScore": 84.4,
"grammarScore": 90.6,
"topicScore": 73.0,
"emphasis": 2,
"wordAdvice": "单词拼写极少出现错误，或无错误；词汇很丰富，用词很准确，超厉害！",
"structureScore": 83.8,
"structureAdvice": "能够灵活使用多种衔接手段，超厉害！"}
*/
type MajorScore struct {
	GrammarAdvice   string  `json:"grammarAdvice"`
	WordScore       float32 `json:"wordScore"`
	GrammarScore    float32 `json:"grammarScore"`
	TopicScore      float32 `json:"topicScore"`
	WordAdvice      string  `json:"wordAdvice"`
	StructureScore  float32 `json:"structureScore"`
	StructureAdvice string  `json:"structureAdvice"`
}
type SentFeedback struct {
	RawSent               string          `json:"rawSent"`
	ParaId                int             `json:"paraId"`
	SentId                int             `json:"sentId"`
	ErrorPosInfos         []ErrorPosInfos `json:"errorPosInfos"` // assuming this is an empty array
	SentFeedback          string          `json:"sentFeedback"`
	SentStartPos          int             `json:"sentStartPos"`
	CorrectedSent         string          `json:"correctedSent"`
	RawSegSent            string          `json:"rawSegSent"`
	IsContainGrammarError bool            `json:"isContainGrammarError"`
	IsContainTypoError    bool            `json:"isContainTypoError"`
	SentScore             float64         `json:"sentScore"`
	IsValidLangSent       bool            `json:"isValidLangSent"`
}
type ErrorPosInfos struct {
	StartPos       int    `json:"startPos"`
	EndPos         int    `json:"endPos"`
	ErrorTypeTitle string `json:"errorTypeTitle"`
	ErrBaseInfo    string `json:"errBaseInfo"`
	KnowledgeExp   string `json:"knowledgeExp"`
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
	9: "gre",
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
