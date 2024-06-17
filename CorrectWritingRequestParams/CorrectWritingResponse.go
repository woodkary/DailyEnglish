package correctWriting

import (
	"encoding/json"
	"fmt"
)

// 解析CorrectWritingResult的结构体
type CorrectWritingResult struct {
	RequestId string      `json:"RequestId"`
	ErrorCode string      `json:"errorCode"`
	Result    EssayResult `json:"Result"`
}

type EssayResult struct {
	RawEssay         string          `json:"rawEssay"`
	PiGaiReqTextType int             `json:"piGaiReqTextType"`
	SentNum          int             `json:"sentNum"`
	StLevelCode      int             `json:"stLevelCode"`
	UniqueKey        string          `json:"uniqueKey"`
	EssayAdvice      string          `json:"essayAdvice"`
	Title            string          `json:""`
	TotalScore       float64         `json:"totalScore"`
	WriteType        int             `json:"writeType"`
	EssayLangName    string          `json:"essayLangName"`
	MajorScore       MajorScore      `json:"majorScore"`
	AllFeatureScore  AllFeatureScore `json:"allFeatureScore"`
	ParaNum          int             `json:"paraNum"`
	EssayFeedback    EssayFeedback   `json:"essayFeedback"`
	WordNum          int             `json:"wordNum"`
	FullScore        float64         `json:"fullScore"`
	ArticleFormCode  int             `json:"articleFormCode"`
	ShowStat         []ShowStat      `json:"showStat"`
	TotalEvaluation  string          `json:"totalEvaluation"`
	StLevel          string          `json:"stLevel"`
	ConjWordNum      int             `json:"conjWordNum"`
	WriteModel       int             `json:"writeModel"`
}

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

type AllFeatureScore struct {
	NeuralScore   float64 `json:"neuralScore"`
	Conjunction   float64 `json:"conjunction"`
	Grammar       float64 `json:"grammar"`
	Spelling      float64 `json:"spelling"`
	AdvanceVocab  float64 `json:"advanceVocab"`
	WordNum       float64 `json:"wordNum"`
	Topic         float64 `json:"topic"`
	LexicalSubs   float64 `json:"lexicalSubs"`
	WordDiversity float64 `json:"wordDiversity"`
	SentComplex   float64 `json:"sentComplex"`
	Structure     float64 `json:"structure"`
}

type EssayFeedback struct {
	SentsFeedback []SentenceFeedback `json:"sentsFeedback"`
}

type SentenceFeedback struct {
	RawSent               string        `json:"rawSent"`
	ParaId                int           `json:"paraId"`
	SentId                int           `json:"sentId"`
	ErrorPosInfos         []interface{} `json:"errorPosInfos"`
	SentFeedback          string        `json:"sentFeedback"`
	SentStartPos          int           `json:"sentStartPos"`
	CorrectedSent         string        `json:"correctedSent"`
	RawSegSent            string        `json:"rawSegSent"`
	IsContainGrammarError bool          `json:"isContainGrammarError"`
	IsContainTypoError    bool          `json:"isContainTypoError"`
	SentScore             float64       `json:"sentScore"`
	IsValidLangSent       bool          `json:"isValidLangSent"`
}

type ShowStat struct {
	Item   string        `json:"item"`
	Count  int           `json:"count"`
	Detail []interface{} `json:"detail"`
}

func FormatResult(result *CorrectWritingResult) {
	fmt.Printf("Request ID: %s\n", result.RequestId)
	fmt.Printf("Error Code: %s\n", result.ErrorCode)
	fmt.Printf("Raw Essay: %s\n", result.Result.RawEssay)
	fmt.Printf("Essay Advice: %s\n", result.Result.EssayAdvice)
	fmt.Printf("Total Score: %.2f\n", result.Result.TotalScore)
	fmt.Printf("Major Score:\n")
	fmt.Printf("\tGrammar Advice: %s\n", result.Result.MajorScore.GrammarAdvice)
	fmt.Printf("\tWord Score: %.2f\n", result.Result.MajorScore.WordScore)
	fmt.Printf("\tGrammar Score: %.2f\n", result.Result.MajorScore.GrammarScore)
	fmt.Printf("\tTopic Score: %.2f\n", result.Result.MajorScore.TopicScore)
	fmt.Printf("\tWord Advice: %s\n", result.Result.MajorScore.WordAdvice)
	fmt.Printf("\tStructure Score: %.2f\n", result.Result.MajorScore.StructureScore)
	fmt.Printf("\tStructure Advice: %s\n", result.Result.MajorScore.StructureAdvice)
	fmt.Printf("All Feature Score:\n")
	fmt.Printf("\tNeural Score: %.2f\n", result.Result.AllFeatureScore.NeuralScore)
	fmt.Printf("\tConjunction: %.2f\n", result.Result.AllFeatureScore.Conjunction)
	fmt.Printf("\tGrammar: %.2f\n", result.Result.AllFeatureScore.Grammar)
	fmt.Printf("\tSpelling: %.2f\n", result.Result.AllFeatureScore.Spelling)
	fmt.Printf("\tAdvance Vocab: %.2f\n", result.Result.AllFeatureScore.AdvanceVocab)
	fmt.Printf("\tWord Num: %.2f\n", result.Result.AllFeatureScore.WordNum)
	fmt.Printf("\tTopic: %.2f\n", result.Result.AllFeatureScore.Topic)
	fmt.Printf("\tLexical Subs: %.2f\n", result.Result.AllFeatureScore.LexicalSubs)
	fmt.Printf("\tWord Diversity: %.2f\n", result.Result.AllFeatureScore.WordDiversity)
	fmt.Printf("\tSent Complex: %.2f\n", result.Result.AllFeatureScore.SentComplex)
	fmt.Printf("\tStructure: %.2f\n", result.Result.AllFeatureScore.Structure)
	fmt.Printf("Essay Feedback:\n")
	for _, feedback := range result.Result.EssayFeedback.SentsFeedback {
		fmt.Println()
		fmt.Printf("\tSentence: %s\n", feedback.RawSent)
		fmt.Printf("\tCorrected Sentence: %s\n", feedback.CorrectedSent)
		fmt.Printf("\tSentence Score: %.2f\n", feedback.SentScore)
	}
}

// 定义从 JSON 格式的 []byte 转换为 CorrectWritingResult 对象的函数
func ParseResultFromJSON(jsonData []byte) (*CorrectWritingResult, error) {
	var result CorrectWritingResult
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// 将评价中的essayAdvice，majorScore.grammarAdvice，majorScore.wordAdvice，majorScore.structureAdvice单独抽出来并转化为json格式
func ParseAdviceFromResult(result *CorrectWritingResult) ([]byte, error) {
	grammarAdvice, err := json.Marshal(result.Result.MajorScore.GrammarAdvice)
	if err != nil {
		return nil, err
	}
	wordAdvice, err := json.Marshal(result.Result.MajorScore.WordAdvice)
	if err != nil {
		return nil, err
	}
	structureAdvice, err := json.Marshal(result.Result.MajorScore.StructureAdvice)
	if err != nil {
		return nil, err
	}
	advice, err := json.Marshal(map[string]interface{}{
		"grammarAdvice":   string(grammarAdvice),
		"wordAdvice":      string(wordAdvice),
		"structureAdvice": string(structureAdvice),
	})
	if err != nil {
		return nil, err
	}
	return advice, nil
}
