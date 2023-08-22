package Solution1071

func solution(str1 string, str2 string) string {
	var min_len int
	if len(str1) < len(str2) {
		min_len = len(str1)
	} else {
		min_len = len(str2)
	}
	for i := min_len; i > 0; i-- {
		if len(str1)%i == 0 && len(str2)%i == 0 {
			if str1[:i] == str2[:i] && multiplyString(str1[:i], len(str1)/i) == str1 && multiplyString(str2[:i], len(str2)/i) == str2 {
				return str1[:i]
			}
		}
	}
	return ""
}

func multiplyString(str string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += str
	}
	return result
}
