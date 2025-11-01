package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func EscapeLikeKeyword(keyword string) string {
	keyword = strings.ReplaceAll(keyword, "'", "''")
	keyword = strings.ReplaceAll(keyword, "%", "\\%")
	keyword = strings.ReplaceAll(keyword, "_", "\\_")
	return keyword
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, item) { // 忽略大小写比较
			return true
		}
	}
	return false
}

func RandString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	data := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, data[r.Intn(len(data))])
	}
	return string(result)
}

func GenerateRandomCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < length; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	return code
}

func Stamp2str(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

func JsonEncode(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func JsonDecode(src string, dest interface{}) error {
	return json.Unmarshal([]byte(src), dest)
}
