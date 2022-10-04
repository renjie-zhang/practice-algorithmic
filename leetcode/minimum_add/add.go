package minimum_add

func minAddToMakeValid(s string) (ans int) {
	cnt := 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else if cnt > 0 {
			cnt--
		} else {
			ans++
		}
	}
	return ans + cnt
}
