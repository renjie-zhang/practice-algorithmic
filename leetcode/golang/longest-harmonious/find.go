package longest_harmonious


func findLHS(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c1 := cnt[num+1]; c1 > 0 && c+c1 > ans {
			ans = c + c1
		}
	}
	return
}