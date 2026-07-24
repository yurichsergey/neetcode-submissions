func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    // I think we should calculate two arrays:
    // 1) minimum value from left part from current position
    // 2) maximum value from right part from current position
    minLeft := make([]int, len(prices)-1)
    minLeftPrice := prices[0]
    for i := 1; i < len(prices); i++ {
        minLeftPrice = min(minLeftPrice, prices[i-1])
        minLeft[i-1] = minLeftPrice
    }

    maxRight := make([]int, len(prices)-1)
    maxRightPrice := prices[len(prices)-1]
    for i := len(prices)-2; i >= 0; i-- {
        maxRightPrice = max(maxRightPrice, prices[i])
        maxRight[i] = maxRightPrice
    }

    bestProfit := 0
    for i := 0; i < len(prices)-1; i++ {
        bestProfit = max(bestProfit, maxRight[i]-minLeft[i])
    }
    return bestProfit
}
