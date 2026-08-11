class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        bestProfit = 0
        lowestPrice = prices[0]
        for i = 0; i < len(prices); i++:
            currentPrice = prices[i]
            currentProfit = currentProfit - lowestPrice
            if currentPrice < lowestPrice:
                lowestPrice = currentPrice
            elif currentProfit > bestProfit:
                bestProfit = currentProfit
        return bestPrice
        