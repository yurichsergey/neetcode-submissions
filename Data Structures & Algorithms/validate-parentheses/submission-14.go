func isValid(s string) bool {
    if (len(s) % 2 != 0) {
        return false;
    }
    st := make([]rune, 0, len(s))
    for _, c := range(s) {
        if c == '(' {
            st = append(st, ')')
        } else if c == '[' {
            st = append(st, ']')
        } else if c == '{' {
            st = append(st, '}')
        } else if len(st) == 0 || c != st[len(st) - 1] {
            return false
        } else {
            st = st[:len(st) - 1]
        }
    }
    return len(st) == 0;
}
