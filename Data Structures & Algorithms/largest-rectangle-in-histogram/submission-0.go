func largestRectangleArea(heights []int) int {
	bars := make([]int, len(heights)+2)
	copy(bars[1:],heights)
	
	stack := []int{}
	maxArea := 0

	for i, h := range bars {
		for len(stack) > 0 && h < bars[stack[len(stack)-1]] {
			height := bars[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] -1
			area :=  height * width
			maxArea = max(maxArea, area)
		}
		stack = append(stack, i)
	}

	return maxArea
}