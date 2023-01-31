package check

import (
	"fmt"
	"testing"
)

func TestCheckXMatrix(t *testing.T) {
	var data = [][]int{{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2}}
	// var data = [][]int{{0, 0, 0, 0, 1}, {0, 4, 0, 1, 0}, {0, 0, 5, 0, 0}, {0, 5, 0, 2, 0}, {4, 0, 0, 0, 2}}
	fmt.Print(CheckXMatrix(data))
}
