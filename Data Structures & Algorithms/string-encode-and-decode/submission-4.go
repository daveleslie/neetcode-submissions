type Solution struct{}

func (s *Solution) Encode(strs []string) string {
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
	for remaining != "" {
		idx := strings.IndexByte(remaining, '#')
		strLength, _ := strconv.Atoi(remaining[:idx])
		rest := remaining[idx+1:]
		result = append(result, rest[:strLength])
		remaining = rest[strLength:]

	}

	return result
}
