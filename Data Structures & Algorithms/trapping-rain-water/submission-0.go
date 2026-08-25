func trap(height []int) int {
	var amount int
	// Precalculate prefixes and suffixes for each index
	prefix := make([]int, len(height))
	suffix := make([]int, len(height))
	var lastPrefix, lastSuffix int
	for p, s := 0, len(height)-1; p < len(height); p, s = p+1, s-1 {
		prefix[p] = lastPrefix
		if height[p] > lastPrefix {
			lastPrefix = height[p]
		}
		suffix[s] = lastSuffix
		if height[s] > lastSuffix {
			lastSuffix = height[s]
		}
	}

	// Calculate water held at each index
	for i, h := range height {
		indexAmount := min(prefix[i], suffix[i]) - h
		if indexAmount > 0 {
			amount += indexAmount
		}
	}

	return amount
}
