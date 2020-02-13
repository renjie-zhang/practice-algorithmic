package decryptstring

/**
题目链接：https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
*/

func freqAlphabets(s string) string {
	var answer string
	length := len(s)
	for i := 0; i < length; i++ {
		if i+2 < length && s[i+2] == '#' {
			answer += string((rune(s[i])-'0')*10 + (rune(s[i+1]) - '1' + 'a'))
			i += 2
		} else {
			answer += string(s[i] - '1' + 'a')
		}
	}
	return answer
}
