func isValid(s string) bool {
	closeToOpen := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		expectedBracket, isClosing := closeToOpen[c]
		if !isClosing {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != expectedBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}