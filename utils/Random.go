package utils

import "math/rand"

//生成随机字符串, n为字符串长度
func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

//生成随机数字串
func RandomNcode(n int) string {
	const letters = "1234567890"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
