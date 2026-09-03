func getMaxFreq(counts *[26]int) int {
	var result int
	for _, count := range counts {
		result = max(count, result)
	}
	return result
}

func characterReplacement(s string, k int) int {
	result := 0
	left := 0
	maxFreq := 0
	var counts [26]int

	// The invariant is that at the end of every iteration,
	// the window given by s[left..right] is at most as long
	// as the result,  and counts carries the frequencies of 
	// each character inside the window
	for right := 0; right < len(s); right++ {
		// Extend the window by admitting the next s[right] and incrementing counts
		counts[s[right]-'A']++
		maxFreq = max(counts[s[right]-'A'],  maxFreq)

		// shrink the window if it is invalid. The cost to make a window valid
		// is given by cost = len(window) - maxFreq. The cost must be less than or 
		// equal to the number of substitutions allowed, k. Therefore a window is valid if cost <= k
		if (right-left+1) - maxFreq > k {
			counts[s[left]-'A']--
			left++
			maxFreq = getMaxFreq(&counts)
		}

		result = max(right-left+1, result)
	}
	return result
}
