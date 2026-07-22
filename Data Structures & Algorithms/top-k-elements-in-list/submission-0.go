func topKFrequent(nums []int, k int) []int {

    freq := map[int]int{}
    for _, i := range nums {
        freq[i] += 1
    }

    keys := make([]int, 0, len(freq))
    for i := range freq {
        keys = append(keys, i)
    }

    sort.Slice(keys, func(i, j int) bool {
        return freq[keys[i]] > freq[keys[j]]
    })

    fmt.Println("print freq")
    for i := range(freq) {
        fmt.Println("%i - %i", i, freq[i])
    }
    fmt.Println("print keys")
    for i := range(keys) {
        fmt.Println("%i", keys[i])
    }

    res := make([]int, 0, k)
    for i := 0; i < k; i++ {
        res = append(res, keys[i])
    }
    return res
}
