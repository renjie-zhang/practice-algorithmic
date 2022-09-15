package jump_game

func CanJump(nums []int) bool {
	if nums == nil {
		return false
	}
	var le = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= le {
			le = i
		}
	}
	if le == 0 {
		return true
	}
	return false
}
