class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s: string, t: string): boolean {
        if (s.length !== t.length) {
            return false
        }
        const sObj: Record<string,number> = {}; 
        const tObj: Record<string, number> = {};
        for (let i = 0; i < s.length; i++) {
            const sVal = s[i]
            const tVal = t[i]
            sObj[sVal] === undefined ? sObj[sVal] = 1 : sObj[sVal] += 1;
            tObj[tVal] === undefined ? tObj[tVal] = 1 : tObj[tVal] += 1;
        }

        for (const key of Object.keys(sObj)) {
            if (sObj[key] !== tObj[key]) {
                return false
            }
        }
        return true
    }
}