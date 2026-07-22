func isValid(s string) bool {   
    runes := map[rune]rune{')': '(', ']': '[', '}': '{'}
    stack := []rune{}
    result := true
    for _, r := range s {
        if opened, ok := runes[r]; ok {
            if len(stack) == 0 || stack[len(stack) - 1] != opened {
                result = false
                break
            }
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, r)
        }
    }
    return result && len(stack) == 0
}
