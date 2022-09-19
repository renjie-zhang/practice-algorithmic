package sort_array

import "sort"

func FrequencySort(nums []int) []int {
	var tempMap = make(map[int]int)
	for _, k := range nums {
		tempMap[k]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return tempMap[a] < tempMap[b] || tempMap[a] == tempMap[b] && a > b
	})
	return nums
}
