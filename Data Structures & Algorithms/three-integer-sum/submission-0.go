func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// freq := make(map[int]int)
	// for _, num := range(nums) {
	// 	freq[num]++
	// }

	res := map[[3]int]struct{}{}
	for i := range(nums) {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] + nums[j] + nums[k] == 0 {
					res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}

	resTriplet := [][]int{}
	for triplet := range res {
		resTriplet = append(resTriplet, []int{triplet[0], triplet[1], triplet[2]})
	}
	return resTriplet
}
