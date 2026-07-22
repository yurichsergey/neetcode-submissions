func isPalindrome(s string) bool {
    clearedString := regexp.MustCompile("[^a-zA-Z0-9]").ReplaceAllString(s, "")
    clearedString = strings.ToLower(clearedString)
    runes := []rune(clearedString)
    n := len(runes)
    for i := 0; i < n / 2; i++ {
        if runes[i] != runes[n - 1 - i] {
            return false
        }
    }
    return true
}
