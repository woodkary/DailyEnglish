package test

import (
	utils "DailyEnglish/utils"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	//接收一个控制台输入的字符串，控制台输出一行加密后的字符串和解密后的字符串
	//控制台接收字符串
	str := "123"
	//加密
	encryptStr := utils.AesEncrypt(str, "DailyEnglish_sec")
	fmt.Println("加密后的字符串：", encryptStr)
	//解密
	decryptStr := utils.AesDecrypt(encryptStr, "DailyEnglish_sec")
	fmt.Println("解密后的字符串：", decryptStr)
}
