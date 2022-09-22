package check_array

import (
	"fmt"
	"testing"
)

func TestCanFormArray(t *testing.T) {
	var te = []int{91, 2, 4, 64, 5, 78, 12, 9}
	var te2 = [][]int{{78, 12, 3}, {4, 64, 5}, {91, 2}}
	fmt.Println(CanFormArray(te, te2))
}
