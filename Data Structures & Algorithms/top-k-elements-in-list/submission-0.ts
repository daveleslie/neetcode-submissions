class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums: number[], k: number): number[] {
        const map = new Map<number, number>()
        for (const num of nums) {
            const frequency = map.get(num) ?? 0
            map.set(num, frequency + 1)
                        
        }

        const frequenciesSorted = [...map.entries()].sort((a, b) => {
            return b[1] - a[1]
        }).map(f => f[0])

        

        return frequenciesSorted.slice(0,k)
    }
}
