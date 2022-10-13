package max_chunks

// https://leetcode.cn/problems/max-chunks-to-make-sorted/description/
func maxChunksToSorted(arr []int) (ans int) {
	mx := 0
	for i, x := range arr {
		if x > mx {
			mx = x
		}
		if mx == i {
			ans++
		}
	}
	return
}
