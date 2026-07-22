func longestConsecutive(nums []int) int {
    mp := make(map[int]int)
    res := 0
    for _, n := range nums {
        _, ok := mp[n]
        if ok {
            continue
        }
        length := mp[n-1]+mp[n+1]+1
        mp[n] = length
        // update left and right boundaries
        mp[n  - mp[n-1]] = length
        mp[n+mp[n+1]] = length
        if length > res{
            res = length
        }
    }
    return res
}
