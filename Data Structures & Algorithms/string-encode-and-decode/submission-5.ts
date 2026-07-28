class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs: string[]): string {
		let encoding = ''
		for (const str of strs) {
			encoding += str.length +  "#" + str
		}

		return encoding
	}

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str: string): string[] {
        if (str.length === 0) {
            return []
        }
        let result: string[] = []
        let remaining = str
        let idx = 0
        while (remaining.length !== 0) {
            idx = remaining.indexOf("#")
            const strLength = parseInt(remaining.slice(0, idx))
            const rest = remaining.slice(idx+1)
            result.push(rest.slice(0, strLength))
            remaining = rest.slice(strLength)
        }

        return result
    }
}
