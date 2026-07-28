func groupAnagrams(strs []string) [][]string {
    var result [][]string
    groups := make(map[[26]int][]string)
    for _, s := range strs {
        var counts [26]int
        for _, c := range s {
            charIndex := int(c) - int('a')
            counts[charIndex]++
        }

        groups[counts] = append(groups[counts], s)
    }

    for _, key := range groups {
        result = append(result, key)
    }

    return result
}