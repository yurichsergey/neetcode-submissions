func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	freq := make(map[int]int)
	for _, num := range(nums) {
		freq[num]++
	}
	res := map[[3]int]struct{}{}
	fmt.Printf("feq = %v\n", freq)
	fmt.Printf("nums = %v\n", nums)
	for i := range(nums) {
		freq[nums[i]]--
		fmt.Printf("\niterate by i = %v, nums[i] = %v, freq = %v\n", i, nums[i], freq)
		for j := i + 1; j < len(nums); j++ {
			freq[nums[j]]--
			target := -(nums[i] + nums[j])

			fmt.Printf("%v\n", freq)
			fmt.Printf(
				"nums[i] = %v;  nums[j] = %v; target = %v\n",
				nums[i],
				nums[j],
				target,
			)
			if target < nums[j] {
				fmt.Printf("SKIP target = %v\n", target)
				freq[nums[j]]++
				continue				
			}
			if freqVal, ok := freq[target]; ok && freqVal > 0 {
				fmt.Printf(
					"FOUND = %v, freqVal = %v, ok = %v\n",
					[3]int{nums[i], nums[j], target},
					freqVal,
					ok,
				)
				res[[3]int{nums[i], nums[j], target}] = struct{}{}
			}
			freq[nums[j]]++
		}
	}
	fmt.Printf("\nres = %v\n", res)

	resTriplet := [][]int{}
	for triplet := range res {
		resTriplet = append(resTriplet, []int{triplet[0], triplet[1], triplet[2]})
	}
	return resTriplet
}
