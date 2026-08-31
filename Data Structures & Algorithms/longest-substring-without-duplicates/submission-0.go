func lengthOfLongestSubstring(s string) int {
	var result int 
	seen := make(map[byte]int)

	left := 0

	for right := 0; right < len(s); right++ {
		if index, ok := seen[s[right]]; ok && index >= left {
			left = index + 1
		}
		seen[s[right]] = right
		result = max(result, right+1-left)
	}
	return result
}
