package check_array

import "fmt"

func CanFormArray(arr []int, pieces [][]int) bool {
	var tempMap = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		tempMap[arr[i]] = i
	}
	fmt.Println(tempMap)
	for i := 0; i < len(pieces); i++ {
		if len(pieces[i]) == 1 {
			continue
		}
		for j := 0; j < len(pieces[i])-2; j++ {
			fmt.Printf("j=%d\n", tempMap[pieces[i][j]])
			if _, ok := tempMap[pieces[i][j]]; !ok {
				return false
			}
			fmt.Printf("j+1= %d\n", tempMap[pieces[i][j+1]])
			if _, ok := tempMap[pieces[i][j+1]]; !ok {
				return false
			}
			if tempMap[pieces[i][j]] > tempMap[pieces[i][j+1]] {
				return false
			}
		}
	}
	return true
}
