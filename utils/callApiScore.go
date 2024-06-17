package utils

type CorrectWritingRequestParams struct {
	Q                 []string `json:"q"`     //正文文本，如果是图片，这里是图片的base64编码
	Grade             []string `json:"grade"` //文章等级，可选值：default，elementary，junior，senior，cet4，cet6，graduate_a1，graduate_b1，graduate_b2，graduate_a1
	Title             []string `json:"title"`
	ModelContent      []string `json:"modelContent"`
	IsNeedSynonyms    []string `json:"isNeedSynonyms"`
	CorrectVersion    []string `json:"correctVersion"`
	IsNeedEssayReport []string `json:"isNeedEssayReport"`
}

// func CallApiScore(data []byte, needDecode bool) {
// 	file, err := os.OpenFile
// }
