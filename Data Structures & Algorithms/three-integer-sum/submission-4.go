func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := map[[3]int]struct{}{}
	for i := range nums {
		target := nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[l] + nums[r]
			if sum > -target {
				r--
			} else if sum < -target {
				l++
			} else {
				res[[3]int{nums[i], nums[l], nums[r]}] = struct{}{}
				l++
				r--
			}
		}
	}
	resList := [][]int{}
	for triplet := range res {
		resList = append(resList, triplet[:])
	}
	return resList
}
