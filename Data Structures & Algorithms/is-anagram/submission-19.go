func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap := make(map[rune]int)
    tMap := make(map[rune]int)

    for i, val := range s {
        sMap[val] += 1
        tMap[rune(t[i])] += 1
    }

    for i, sVal := range sMap {
        tVal, exists := tMap[i]
        if !exists || tVal != sVal {
            return false
        }
    }

    return true

    
}
