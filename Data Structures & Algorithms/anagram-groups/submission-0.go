func groupAnagrams(strs []string) [][]string {

	createFrequencyKey := func(s string) string {
		m := map[rune]int{}
		for _, r := range s {
			m[r] += 1
		}

		runes := make([]rune, 0, len(m))
		for r := range m {
			runes = append(runes, r)
		}

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		res := ""
		for _, r := range runes {
			res += string(r) + strconv.Itoa(m[r])
		}
		return res
	}
                        
    m := map[string][]string{}
    for _, s := range strs {
        k := createFrequencyKey(s)
        _, ok := m[k]
        if (!ok) {
            m[k] = []string{}
        }
        m[k] = append(m[k], s)
    }
    
    res := [][]string{}
    for _, v := range m {
        res = append(res, v)
    }
    return res
}

