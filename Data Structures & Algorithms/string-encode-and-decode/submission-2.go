type Solution struct{}

func (s *Solution) Encode(strs []string) string {
		if len(strs) == 0 {
		return ""
	}

	// Encode the lengths of the strings
	sizes := strings.Builder{}
	for i, str := range strs {
		sizes.WriteString(strconv.Itoa(len(str)))
		if i != len(strs)-1 {
			sizes.WriteString(",")
		}
	}
	sizes.WriteString("#")

	// Append the strings themselves and return the final encoded string
	return sizes.String() + strings.Join(strs, "")

}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	// get the string lengths from the encoded string
	parts := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(parts[0], ",")

	var result []string
	start := 0
	for _, sizeStr := range sizes {
		size, _ := strconv.Atoi(sizeStr)
		result = append(result, parts[1][start:start+size])
		start += size
	}

	return result

}
