func productExceptSelf(nums []int) []int {
	right := make([]int, len(nums)-1)
	right[len(nums)-2] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		right[i-1] = right[i] * nums[i]
	}

	left := make([]int, len(nums)-1)
	left[0] = nums[0]
	for i := 1; i < len(nums)-1; i++ {
		left[i] = left[i-1] * nums[i]
	}

	// nums    = [   2,   3,   4,   5,]
	// nums_i  = [   0,   1,   2,  ..,]
	// nums_i  = [ i-1,   i, i+2,   3,]
	//
	// right   = [  60,  20,   5,     ]
	// right_i = [   0,   1,   2,     ]
	// left    = [        2,   6,  24,]
	// left_i  = [        0,   1,   2,]
	res := make([]int, len(nums))
	res[0] = right[0]
	res[len(nums)-1] = left[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		res[i] = right[i] * left[i-1]
	}

	leftJson, _ := json.Marshal(left)
	fmt.Println(string(leftJson))
	return res
}
