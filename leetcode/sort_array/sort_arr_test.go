package sort_array

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	var te = []int{1, 1, 2, 2, 2, 3}
	fmt.Println(FrequencySort(te))
}
