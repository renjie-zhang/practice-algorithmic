package build_array

// https://leetcode.cn/problems/build-an-array-with-stack-operations/
func BuildArray(target []int, n int) []string {
	var index = 0
	var result []string
	var length = len(target)
	for i := 1; i <= n; i++ {
		if i == target[index] {
			result = append(result, "Push")
			index++
			if length == index {
				break
			}
		} else {
			result = append(result, "Push")
			result = append(result, "Pop")
		}
	}
	return result
}
