package utils

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

func EncryptIC(teamID int, key int) string {
	// 将 teamID 和 key 进行异或
	obfuscated := teamID ^ key
	// 将结果转换为字符串
	obfuscatedStr := strconv.Itoa(obfuscated)
	// 通过重复 key 和 obfuscatedStr 创建一个较长的字符串
	longStr := obfuscatedStr + strconv.Itoa(key)
	fmt.Println("Long string:", longStr)
	fmt.Println("Long string bytes:", []byte(longStr))
	// 使用 Base64 编码来混淆字符串
	encoded := base64.StdEncoding.EncodeToString([]byte(longStr))
	// 确保输出为 8 位，如果不足则填充，如果超出则截断
	if len(encoded) < 8 {
		encoded = encoded + strings.Repeat("A", 8-len(encoded))
	} else if len(encoded) > 8 {
		encoded = encoded[:8]
	}
	return encoded
}

func DecryptIC(encrypted string, key int) (int, error) {
	// 解码 Base64 字符串
	decodedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return 0, err
	}
	//打印解码后的字符串
	fmt.Println("Decoded bytes:", decodedBytes)
	decodedStr := string(decodedBytes)

	// 移除附加的 key 字符串部分
	keyStr := strconv.Itoa(key)
	if strings.HasSuffix(decodedStr, keyStr) {
		obfuscatedStr := strings.TrimSuffix(decodedStr, keyStr)
		// 将字符串转换回整数
		obfuscated, err := strconv.Atoi(obfuscatedStr)
		if err != nil {
			return 0, err
		}
		// 使用与加密相同的异或操作解密
		teamID := obfuscated ^ key
		return teamID, nil
	}
	return 0, fmt.Errorf("decryption failed")
}

func TestICD() {
	teamID := 12345
	key := 114514
	encrypted := EncryptIC(teamID, key)
	fmt.Println("Encrypted string:", encrypted)
	decryptedTeamID, _ := DecryptIC(encrypted, key)
	fmt.Println("Decrypted team ID:", decryptedTeamID)
}
