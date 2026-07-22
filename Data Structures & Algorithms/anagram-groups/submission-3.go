func groupAnagrams(strs []string) [][]string {

	createFrequencyKey := func(s string) string {
		m := [26]int{}
		for _, r := range s {
			m[r-'a'] += 1
		}

		var res strings.Builder
		for i, count := range m {
			res.WriteByte(byte('a' + i))
			res.WriteString(strconv.Itoa(count))
		}
		return res.String()
	}
                            
    m := map[string][]string{}
    for _, s := range strs {
        k := createFrequencyKey(s)
        m[k] = append(m[k], s)
    }
    
    res := make([][]string, 0, len(m))
    for _, v := range m {
        res = append(res, v)
    }
    return res
}

