func minWindow(s string, t string) string {
    if t == "" {
		return ""
	}
	var result string
	resLen := math.MaxInt

	countsT := make(map[byte]int)
	for i := range len(t) {
		countsT[t[i]]++
	}
	countsWindow := make(map[byte]int)

	have, need := 0, len(countsT)

	// The invariant is that at the start of every iteration
	for l, r := 0, 0; r < len(s); r++ {
		countsWindow[s[r]]++
		// If new right edge character is in t, and the frequency of that character in t matches that inside window,
		// then update have 
		if countsT[s[r]] > 0 && countsWindow[s[r]] == countsT[s[r]] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				result = s[l:r+1]
				resLen = r - l + 1
			}
			countsWindow[s[l]]--
			if countsT[s[l]] > 0 && countsT[s[l]] > countsWindow[s[l]] {
				have--
			}
			l++

		}
	}

	return result
}
