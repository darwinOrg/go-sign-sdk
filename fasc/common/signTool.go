package common

import (
	"sort"
	"strings"
)

// GetSignByMap 计算签名参数
func GetSignByMap(headMap map[string]string, timestamp string, appSecret string) string {
	signContent := SortMapForSign(headMap)
	signText := Sha256(signContent)
	secretSigningByte := HmacSha256byte(timestamp, appSecret) //临时签名密钥
	sign := HmacSha256(signText, secretSigningByte)
	return strings.ToLower(sign)
}

// SortMapForSign 将map key进行排序拼接
func SortMapForSign(headMap map[string]string) string {
	var keys []string
	for k := range headMap {
		keys = append(keys, k)
	}
	var headStr []string
	sort.Strings(keys)
	// To perform the opertion you want
	for _, k := range keys {
		if len(headMap[k]) != 0 {
			headStr = append(headStr, k+"="+headMap[k])
		}
	}
	signContent := strings.Join(headStr, "&")
	return signContent
}
