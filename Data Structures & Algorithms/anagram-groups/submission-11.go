import (
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    var result [][]string
    groups := make(map[string][]string)
    for _, s := range strs {
        sorted := sortString(s)
        groups[sorted] = append(groups[sorted], s)
    }

    for _, key := range groups {
        result = append(result, key)
    }

    return result
}

func sortString(s string) string {
    runes := []rune(s)
    slices.Sort(runes)
    return string(runes)
}
