func maxProfit(prices []int) int {
    bestProfit := 0
    lowestPrice := prices[0]
    for i := 1; i < len(prices); i++ {
        currentPrice := prices[i]
        currentProfit := currentPrice - lowestPrice
        if currentPrice < lowestPrice {
            lowestPrice = currentPrice
        } else if currentProfit > bestProfit {
            bestProfit = currentProfit
        }
    }
    return bestProfit
}
