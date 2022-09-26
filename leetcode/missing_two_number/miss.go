package missing_two_number

func missingTwo(nums []int) []int {
	xorSum := 0
	n := len(nums) + 2
	for _, num := range nums {
		xorSum ^= num
	}
	for i := 1; i <= n; i++ {
		xorSum ^= i
	}
	lsb := xorSum & -xorSum
	type1, type2 := 0, 0
	for _, num := range nums {
		if num&lsb > 0 {
			type1 ^= num
		} else {
			type2 ^= num
		}
	}
	for i := 1; i <= n; i++ {
		if i&lsb > 0 {
			type1 ^= i
		} else {
			type2 ^= i
		}
	}
	return []int{type1, type2}
}
