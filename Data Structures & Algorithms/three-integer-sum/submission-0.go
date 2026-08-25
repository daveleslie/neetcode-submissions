import (
	"slices"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums)
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}
		remaining := nums[i+1:]
		left, right := 0, len(remaining)-1
		target := -num
		for left < right {
			sum := remaining[left] + remaining[right]
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				result = append(result, []int{num, remaining[left], remaining[right]})
				for left++; left < right && remaining[left] == remaining[left-1]; left++ {
				}
			}
		}
	}
	return result
}
