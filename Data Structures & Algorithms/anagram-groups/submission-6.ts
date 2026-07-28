class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const result: string[][] = []
        const map = new Map<string, number>()
        // iterate over strs
        for (const [idx, elem] of strs.entries()) {
            
            const anagram = elem.split("").sort().join("")
            const anagramIndex = map.get(anagram)
            if (anagramIndex !== undefined) {
                result[anagramIndex].push(elem)
                continue
            }

            result.push([elem])
            map.set(anagram, result.length-1)
            
        }

        return result
    }
}
