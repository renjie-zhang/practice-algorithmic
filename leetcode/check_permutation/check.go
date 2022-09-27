package check_permutation

func CheckPermutation(s1 string, s2 string) bool {
	var length1 = len(s1)
	var length2 = len(s2)
	if length1 == 0 && length2 == 0 {
		return true
	}
	if length1 != length2 {
		return false
	}
	var vMap1 = make(map[uint8]int)
	var vMap2 = make(map[uint8]int)
	for i := 0; i < length1; i++ {
		vMap1[s1[i]]++
		vMap2[s2[i]]++
	}
	for k, v := range vMap1 {
		if te, ok := vMap2[k]; ok {
			if v != te {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
