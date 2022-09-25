package rorated_digits

import "strconv"

var check = [10]int{0, 0, 1, -1, -1, 1, 1, -1, 0, 1}

func rotatedDigits(n int) int {
	digits := strconv.Itoa(n)
	memo := [5][2][2]int{}
	for i := 0; i < 5; i++ {
		memo[i] = [2][2]int{{-1, -1}, {-1, -1}}
	}
	var dfs func(int, bool, bool) int
	dfs = func(pos int, bound, diff bool) (res int) {
		if pos == len(digits) {
			return bool2int(diff)
		}
		ptr := &memo[pos][bool2int(bound)][bool2int(diff)]
		if *ptr != -1 {
			return *ptr
		}
		lim := 9
		if bound {
			lim = int(digits[pos] - '0')
		}
		for i := 0; i <= lim; i++ {
			if check[i] != -1 {
				res += dfs(pos+1, bound && i == int(digits[pos]-'0'), diff || check[i] == 1)
			}
		}
		*ptr = res
		return
	}
	return dfs(0, true, false)
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
