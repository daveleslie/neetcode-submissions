func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{})
    for _, val := range nums {
        _, exists := set[val]
        if exists {
            return true
        }
        set[val] = struct{}{}
    }
    return false
}
