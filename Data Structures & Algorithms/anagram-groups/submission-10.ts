class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const map = new Map<string, string[]>()
        // iterate over strs
        for (const elem of strs) {
            
            const key = elem.split("").sort().join("")
            if (!map.has(key)) {
                map.set(key, [])
            }

            map.get(key).push(elem)
        }

        return Array.from(map.values())
    }
}
