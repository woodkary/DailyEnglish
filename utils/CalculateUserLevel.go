package utils

import (
	"strconv"
	"strings"
)

func CalculateUserLevel(ScoresInExam string) []int {
	var AlevelNum int = 0
	var BlevelNum int = 0
	var ClevelNum int = 0
	var DlevelNum int = 0
	var ElevelNum int = 0
	var FlevelNum int = 0
	datas := strings.Split(ScoresInExam, "-")
	for _, data := range datas {
		score, _ := strconv.Atoi(data)
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
