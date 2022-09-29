package string_rotation

func isFlipedString(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	if n == 0 {
		return true
	}
next:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s1[(i+j)%n] != s2[j] {
				continue next
			}
		}
		return true
	}
	return false
}
