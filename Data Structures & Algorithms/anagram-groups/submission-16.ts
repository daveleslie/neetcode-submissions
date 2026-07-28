class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const map = new Map<string, string[]>()
        const offset = "a".charCodeAt(0)
        // iterate over strs
        for (const str of strs) {
            const countMap = Array(26).fill(0)
            for (const char of str) {
                const charCode = char.toLowerCase().charCodeAt(0)
                countMap[charCode - offset]++
            }
            const key = countMap.join("#")
            if (!map.has(key)) {
                map.set(key, [])
            } 
            map.get(key).push(str)
            
        }

        return Array.from(map.values())
    }
}
