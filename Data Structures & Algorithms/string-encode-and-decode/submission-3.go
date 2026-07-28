type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Encode the lengths of the strings
	encoding := strings.Builder{}
	for _, str := range strs {
		encoding.WriteString(strconv.Itoa(len(str)))
		encoding.WriteString("#")
		encoding.WriteString(str)
	}

	// Append the strings themselves and return the final encoded string
	return encoding.String()

}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	var result []string

	// get the string lengths from the encoded string
	remaining := encoded
	for {
		parts := strings.SplitN(remaining, "#", 2)
		if parts[0] == "" {
			break
		}
		strLength, _ := strconv.Atoi(parts[0])
		result = append(result, parts[1][0:strLength])
		remaining = parts[1][strLength:]

	}

	return result
}
