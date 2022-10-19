package unable_eat_lunch

// https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/

func CountStudents(students []int, sandwiches []int) int {
	for {
		if sandwiches[0] != students[0] {
			var te = students[0]
			students = students[1:]
			students = append(students, te)
		} else {
			students = students[1:]
			sandwiches = sandwiches[1:]
		}
		if len(sandwiches) == 0 {
			break
		}
	}
	return len(students)
}
