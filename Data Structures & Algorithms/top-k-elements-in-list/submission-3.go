func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
	for _, i := range nums {
		freq[i] += 1
	}

	revert := make([][]int, len(nums))

	for num, freq := range freq {
		revert[freq-1] = append(revert[freq-1], num)
	}

	res := make([]int, 0, k)
	for i := len(nums); i >= 0; i-- {
		if len(revert[i-1]) == 0 {
			continue
		}
		for _, num := range revert[i-1] {
			res = append(res, num)
			if len(res) >= k {
				return res
			}
		}
	}

	return res
}
