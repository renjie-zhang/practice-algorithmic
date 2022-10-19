package unable_eat_lunch

import (
	"fmt"
	"testing"
)

func TestCountStudents(t *testing.T) {
	var stu = []int{1, 1, 0, 0}
	var san = []int{0, 1, 0, 1}
	fmt.Println(CountStudents(stu, san))
}
