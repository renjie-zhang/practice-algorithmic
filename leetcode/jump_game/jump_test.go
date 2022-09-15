package jump_game

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	var temp = []int{3, 2, 1, 0, 4}
	fmt.Println(CanJump(temp))
}
