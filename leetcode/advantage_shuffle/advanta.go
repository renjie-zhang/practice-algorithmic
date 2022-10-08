package advantage_shuffle

import (
	"fmt"
	"sort"
)

func AdvantageCount(nums1 []int, nums2 []int) []int {
	var length = len(nums1)
	var result = make([]int, length)
	var ids = make([]int, length)
	sort.Ints(nums1)
	for i := 0; i < len(ids); i++ {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool { return nums2[ids[i]] < nums2[ids[j]] })
	left, right := 0, length-1
	for _, x := range nums1 {
		if x > nums2[ids[left]] {
			result[ids[left]] = x
			left++
		} else {
			result[ids[right]] = x
			right--
		}
	}
	fmt.Println(result)
	return result
}
