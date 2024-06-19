package test

import (
	utils "DailyEnglish/utils"
	"fmt"
	"testing"
)

func TestImageCorrect(t *testing.T) {
	Base64IMG, err := utils.ReadFileAsBase64("D:/屏幕截图 2024-06-17 133710.png")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	result := utils.CallApiScore(Base64IMG, 6, "A Proposal to the School Library")
	// 打印返回结果
	fmt.Print(result)
}
