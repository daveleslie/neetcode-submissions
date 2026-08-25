func maxArea(heights []int) int {
	var maxArea int
	left, right := 0, len(heights)-1
	for left < right {
		height := min(heights[left], heights[right])
		width := right - left
		area := height * width
		maxArea = max(maxArea, area)
		if height == heights[left] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
