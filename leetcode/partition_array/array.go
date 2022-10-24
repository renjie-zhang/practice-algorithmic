package partition_array

// https://leetcode.cn/problems/partition-array-into-disjoint-intervals/

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax, leftPos, curMax := nums[0], 0, nums[0]
	for i := 1; i < n-1; i++ {
		curMax = max(curMax, nums[i])
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

