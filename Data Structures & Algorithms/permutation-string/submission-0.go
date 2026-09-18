func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    var targetCounts [26]int
    var windowCounts [26]int

    for i := range len(s1) {
        targetCounts[s1[i]-'a']++
        windowCounts[s2[i]-'a']++
    }

    var matches int
    for i := range 26 {
        if targetCounts[i] == windowCounts[i] {
            matches++
        }
    }

    for r := len(s1); r < len(s2); r++ {
        if matches == 26 {
            return true
        }

        incoming := s2[r]
        windowCounts[incoming-'a']++

        switch windowCounts[incoming-'a'] {
        case targetCounts[incoming-'a']:
            matches++
        case targetCounts[incoming-'a'] + 1:
            matches--
        }

        outgoing := s2[r-len(s1)]
        windowCounts[outgoing-'a']--

        switch windowCounts[outgoing-'a'] {
        case targetCounts[outgoing-'a']:
            matches++
        case targetCounts[outgoing-'a'] - 1:
            matches--
        }
    }

    return matches == 26
}