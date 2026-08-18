func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 'a' - 'A'
	}
	return c
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		// Set left pointer - skip every non-alphanumeric char
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Set right pointer
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		// Check if left != right
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
