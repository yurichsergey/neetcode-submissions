func longestConsecutive(nums []int) int {

    uniq := map[int]struct{}{} 
    for _, i := range nums {
        uniq[i] = struct{}{}
    }

    maxLen := 0
    for _, num := range nums {
        curLen := 1

        for i := 1; true; i++ {
            _, ok := uniq[num + i]
            if (!ok) {
                break
            }
            curLen++
        }

        if curLen > maxLen {
            maxLen = curLen
        }
    }
    return maxLen
}
