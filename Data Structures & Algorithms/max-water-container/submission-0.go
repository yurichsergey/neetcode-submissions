func maxArea(heights []int) int {
	maxArea := 0
	for i := range heights {
		for j := i+1; j < len(heights); j++ {
			currentArea := min(heights[i], heights[j]) * (j - i)
			if currentArea > maxArea {
				maxArea = currentArea
			}
		}
	}
	return maxArea
}
