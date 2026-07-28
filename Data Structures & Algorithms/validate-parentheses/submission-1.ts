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
            // if it is a closing bracket
            if (bracketMap[c as ClosingBracket]) {
                if (stack.length === 0) {
                    return false
                }
                if (stack[stack.length - 1] !== bracketMap[c as ClosingBracket]) {
                    return false
                } 

                if (stack[stack.length - 1] === bracketMap[c as ClosingBracket]) {
                    stack.pop()
                    continue
                }
                stack.push(c)
            } else {
                stack.push(c)
            }


        }

        if (stack.length !== 0) {
            return false
        }

        return true   
    }
}