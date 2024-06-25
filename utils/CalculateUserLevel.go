package utils

func CalculateUserLevel(ScoresInExam []int) []int {
	var AlevelNum int = 0
	var BlevelNum int = 0
	var ClevelNum int = 0
	var DlevelNum int = 0
	var ElevelNum int = 0
	var FlevelNum int = 0
	for _, score := range ScoresInExam {
		if score >= 93 {
			AlevelNum++
		} else if score >= 85 {
			BlevelNum++
		} else if score >= 77 {
			ClevelNum++
		} else if score >= 69 {
			DlevelNum++
		} else if score >= 60 {
			ElevelNum++
		} else {
			FlevelNum++
		}
	}
	return []int{AlevelNum, BlevelNum, ClevelNum, DlevelNum, ElevelNum, FlevelNum}
}
