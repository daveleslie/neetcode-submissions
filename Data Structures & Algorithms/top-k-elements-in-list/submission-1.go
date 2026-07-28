func topKFrequent(nums []int, k int) []int {
	fMap := make(map[int]int)

	for _, num := range nums {
		fMap[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range fMap {
		buckets[freq] = append(buckets[freq], num)
	}

	var result []int
	for i := len(buckets) - 1; i > 0; i-- {
		for _, n := range buckets[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}

	return []int{}
}
