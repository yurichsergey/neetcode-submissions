func maxArea(heights []int) int {
	maxArea := 0
	l, r := 0, len(heights) - 1
	for l < r {
		currentArea := min(heights[l], heights[r]) * (r-l)
		if currentArea > maxArea {
			maxArea = currentArea
		} 
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}
	return maxArea
}
