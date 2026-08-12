func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	// Iterate over temperatures
	for i, temp := range temperatures {
		
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			stackTop := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			result[stackTop] = i - stackTop
		}

		stack = append(stack, i)

	}
	return result
}

