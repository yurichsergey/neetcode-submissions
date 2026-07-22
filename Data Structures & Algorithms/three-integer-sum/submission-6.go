func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := range nums {
		target := nums[i]
		l := i + 1
		r := len(nums) - 1

		if target > 0 {
			break
		}
		if i > 0 && target == nums[i-1] {
			continue
		}

		for l < r {
			sum := nums[l] + nums[r]
			if sum > -target {
				r--
			} else if sum < -target {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}
	return res
}
