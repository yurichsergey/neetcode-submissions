func hasDuplicate(nums []int) bool {
    set := map[int]struct{}{}
    for _, v := range nums {
        _, ok := set[v]
        if ok {
            return true
        } else {
            set[v] = struct{}{}
        }
    }
    return false
}
