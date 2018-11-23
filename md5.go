package comm

import (
	"crypto/md5"
	"encoding/hex"
)

//返回一个字符串的ＭＤ５值
func MD5Sum(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
