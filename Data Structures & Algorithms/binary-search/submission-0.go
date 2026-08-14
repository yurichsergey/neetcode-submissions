func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	for l <= r {
		n := (r + l) / 2
		if nums[n] == target {
			return n
		}
		if nums[n] > target {
			r = n-1
		} else {
			l = n+1
		}
	}
	return -1
}
