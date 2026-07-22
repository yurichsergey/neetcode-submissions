func groupAnagrams(strs []string) [][]string {

	createFrequencyKey := func(s string) string {
		m := [26]int{}
		for _, r := range s {
			m[r-'a'] += 1
		}

		res := ""
		for i, count := range m {
			res += string(byte('a'+i)) + strconv.Itoa(count)
		}
		return res
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

