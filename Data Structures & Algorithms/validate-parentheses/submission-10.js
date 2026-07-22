class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {
        if (s.length % 2 !== 0) {
            return false
        }
        const st = []
        for (const c of s) {
            if ([']', '}', ')'].includes(c)) {
                const last = st.length > 0 ? st.pop() : ''
                if ((c === ']' && last !== '[') ||
                    (c === ')' && last !== '(') ||
                    (c === '}' && last !== '{')
                ) {
                    return false
                }
            } else { // here we save all OPENED brackets
                st.push(c)
            }
        }
        return st.length == 0
    }
}
