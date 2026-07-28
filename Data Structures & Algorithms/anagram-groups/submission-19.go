func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    for _, s := range strs {
        var counts [26]int
        for _, c := range s {
            charIndex := int(c) - int('a')
            counts[charIndex]++
        }

        groups[counts] = append(groups[counts], s)
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}