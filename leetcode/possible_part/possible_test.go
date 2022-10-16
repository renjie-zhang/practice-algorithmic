package possible_part

import (
	"fmt"
	"testing"
)

func TestPossibleBipartition(t *testing.T) {
	dislike := [][]int{{1, 2}, {1, 3}, {2, 4}}
	fmt.Println(PossibleBipartition(4, dislike))
}
