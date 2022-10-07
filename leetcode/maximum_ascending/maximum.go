package maximum_ascending

func maxAscendingSum(nums []int) (ans int) {
	for i, n := 0, len(nums); i < n; {
		sum := nums[i]
		for i++; i < n && nums[i] > nums[i-1]; i++ {
			sum += nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
