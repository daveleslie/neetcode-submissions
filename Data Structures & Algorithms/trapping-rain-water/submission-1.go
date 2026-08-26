func trap(height []int) int {
	var amount int
	
	left, right := 0, len(height)-1
	leftMax, rightMax := height[0], height[len(height)-1]

	for left < right {
		if height[left] < height[right] {
			left++
			leftMax = max(leftMax, height[left])
			amount += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			amount += rightMax - height[right]
		}
	}

	return amount
}
