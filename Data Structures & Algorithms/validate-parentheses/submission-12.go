func isValid(s string) bool {
    if len(s) % 2 != 0 {
        return false
    }

    brackets := map[rune]rune{']':'[', ')':'(', '}':'{'}
    st := []rune{}

    for _, c := range s {
        if openBracket, ok := brackets[c]; ok {
            if len(st) == 0 || openBracket != st[len(st)-1] {
                return false
            }
            st = st[:len(st)-1]
        } else {
            st = append(st, c)
        }
    }
    return len(st) == 0
}
