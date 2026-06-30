func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for i, v := range nums {
        complement := target - v
        if seenIndex, ok := seen[complement]; ok {
            return []int{seenIndex, i}
        }
        seen[v] = i
    }
    return []int{}
}
