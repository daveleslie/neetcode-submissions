class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const map = new Map<string, string[]>()
        // iterate over strs
        for (const elem of strs) {
            
            const anagram = elem.split("").sort().join("")
            const anagramArray = map.get(anagram)
            if (anagramArray !== undefined) {
                anagramArray.push(elem)
                continue
            }

            map.set(anagram, [elem])
        }

        return Array.from(map.values())
    }
}
