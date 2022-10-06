package three_parts

func threeEqualParts(arr []int) []int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}

	partial := sum / 3
	first, second, third, cur := 0, 0, 0, 0
	for i, x := range arr {
		if x == 1 {
			if cur == 0 {
				first = i
			} else if cur == partial {
				second = i
			} else if cur == 2*partial {
				third = i
			}
			cur++
		}
	}

	n := len(arr)
	length := n - third
	if first+length <= second && second+length <= third {
		i := 0
		for third+i < n {
			if arr[first+i] != arr[second+i] || arr[first+i] != arr[third+i] {
				return []int{-1, -1}
			}
			i++
		}
		return []int{first + length - 1, second + length}
	}
	return []int{-1, -1}
}
