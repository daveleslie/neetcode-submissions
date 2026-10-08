func maxSlidingWindow(nums []int, k int) []int {
    var result, dq []int

	// Invariant - After processing i:
	// - dq contains the indexes over the closed interval [i-k+1: i+1], in increasing order
	// - the corresponding nums values of dq are strictly decreasing from front to back
	// - The front of the deque dq[0] always holds the window maximum
	for i := range nums {
		// 1. Expire front of deque if that index has fallen out of window
		if len(dq) > 0 && dq[0] <  i - k + 1 {
			dq = dq[1:]
		}
		// 2. Remove dominated elements from the back of the deque
		for len(dq) > 0 && nums[i] >= nums[dq[len(dq)-1]] {
			dq = dq[:len(dq)-1]
		}
		// 3. Push i to the back of the deque
		dq = append(dq, i)
		// 4. Update result for valid window sizes
		if i >= k-1 {
			result = append(result, nums[dq[0]])
		}
	}
	return result
}
