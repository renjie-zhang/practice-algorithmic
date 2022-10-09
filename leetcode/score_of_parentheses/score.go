package score_of_parentheses

func scoreOfParentheses(s string) int {
	st := []int{0}
	for _, c := range s {
		if c == '(' {
			st = append(st, 0)
		} else {
			v := st[len(st)-1]
			st = st[:len(st)-1]
			st[len(st)-1] += max(2*v, 1)
		}
	}
	return st[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

