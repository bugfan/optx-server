package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptedPassword(str string) string {
	return SHA256Str(str)
}

// sha256验证
func SHA256Str(src string) string {
	h := sha256.New()
	h.Write([]byte(src)) // 需要加密的字符串为
	// fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	return hex.EncodeToString(h.Sum(nil))
}
