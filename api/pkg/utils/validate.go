package utils

import "regexp"

// 手机号校验（中国大陆）

func IsPhoneNumber(phone string) bool {
	// 简单匹配 11 位数字，以 1 开头
	re := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}

// 身份证号校验（支持 18 位，最后一位可能是 X）

func IsIDNumber(id string) bool {
	re := regexp.MustCompile(`^\d{17}[\dXx]$`)
	if !re.MatchString(id) {
		return false
	}
	// 校验最后一位校验码
	return validateIDChecksum(id)
}

// 身份证校验位计算
func validateIDChecksum(id string) bool {
	if len(id) != 18 {
		return false
	}
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCodes := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(id[i]-'0') * weights[i]
	}
	mod := sum % 11
	return checkCodes[mod] == id[17] || (checkCodes[mod] == 'X' && id[17] == 'x')
}
