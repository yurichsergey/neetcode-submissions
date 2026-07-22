func isAnagram(s string, t string) bool {

    buildStat := func (s string) map[rune]int {
        m := map[rune]int{}
        for _, r := range s {
            m[r] += 1
        }
        return m
    }

    ms := buildStat(s)
    mt := buildStat(t)

    for r, i := range ms {
        mt[r] -= i
    }

    res := true
    for _, j := range mt {
        if j != 0 {
            res = false
            break
        }
    }
    return res
}
