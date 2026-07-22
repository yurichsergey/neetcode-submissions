func topKFrequent(nums []int, k int) []int {

    freq := map[int]int{}
    for _, i := range nums {
        freq[i] += 1
    }

    maxKeyByValue := func(m map[int]int) int {
        r, rv := 0, 0
        for k, v := range m {
            if v > rv {
                r = k
                rv = v
            }
        }
        return r
    } 

    res := make([]int, 0, k)
    for i := 0; i < k; i++ {
        maxKey := maxKeyByValue(freq)
        res = append(res, maxKey)
        delete(freq, maxKey)
    }
    return res
}
