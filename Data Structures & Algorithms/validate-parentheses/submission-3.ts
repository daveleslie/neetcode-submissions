type ClosingBracket = ')'|'}'|']'

class Solution {
/**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s: string): boolean {
        const stack: string[] = []
        const bracketMap: Record<ClosingBracket, string> = {
            ')':'(',
            ']':'[',
            '}':'{'
        }
        for (const c of s) {
            const correspondingBracket = bracketMap[c as ClosingBracket]
            // if it is a closing bracket
            if (correspondingBracket !== undefined) {
                if (stack.pop() !== correspondingBracket) {
                    return false
                }
            } else {
                stack.push(c)
            }


        }

        return stack.length === 0   
    }
}