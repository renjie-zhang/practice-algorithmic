package sign_of_product

func arraySign(nums []int) int {
	var result = 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			result = result * -1
		}
		if nums[i] == 0 {
			return 0
		}
		if nums[i] > 0 {
			result = result * 1
		}
	}
	return result
}
