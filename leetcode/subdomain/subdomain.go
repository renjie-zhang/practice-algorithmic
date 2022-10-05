package subdomain

import (
	"fmt"
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {
	var keyMap = make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		// 通过空格去切分
		split := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(split[0])
		// 倒着切分,首先查看字符中存在多长的域名
		tempURL := strings.Split(split[1], ".")
		tempLength := len(tempURL)
		for j := tempLength; j > 0; j-- {
			tempChar := strings.SplitN(split[1], ".", j)
			fmt.Println(tempChar)
			fmt.Println(j)
			keyMap[tempChar[len(tempChar)-1]] += count
		}
	}
	// 构建答案
	var result []string
	for k, v := range keyMap {
		result = append(result, fmt.Sprintf("%d %s", v, k))
	}
	fmt.Println(result)
	return result
}
