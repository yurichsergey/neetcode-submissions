func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		s := numbers[i] + numbers[j]
		if s == target {
			break
		}
		if s < target {
			i++
			continue
		}
		if s > target {
			j--
			continue
		}
	}

	// array is 1-indexed
	i++
	j++

	return []int{i, j}
}
