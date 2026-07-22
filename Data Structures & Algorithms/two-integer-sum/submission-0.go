func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for ki, i := range nums {
        diff := target - i
        kj, ok := m[diff]
        if ok {
            return []int{kj, ki}
        }
        m[i] = ki
    }
    return []int{-1, -1}
}
