package jump_game_2

func Jump(nums []int) int {
	var le = len(nums) - 1
	var steps = 0
	for le > 0 {
		for i := 0; i < le; i++ {
			if i+nums[i] >= le {
				le = i
				steps++
				break
			}
		}
	}
	return steps
}
