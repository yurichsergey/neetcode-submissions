func isPalindrome(s string) bool {
	n := len(s)
	i := 0
	j := n - 1
	re := regexp.MustCompile("[a-zA-Z0-9]")
	for i <= j {
		if !re.MatchString(string(s[i])) {
			i++
			continue
		}
		if !re.MatchString(string(s[j])) {
			j--
			continue
		}
		//print(string(s[i]) + " -- " + string(s[j]) + "\n")
		if strings.ToLower(string(s[i])) == strings.ToLower(string(s[j])) {
			i++
			j--
			continue
		}
		return false
	}
	return true
}
