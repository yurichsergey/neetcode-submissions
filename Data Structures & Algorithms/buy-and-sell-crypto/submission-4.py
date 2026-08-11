class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        bestProfit = 0
        lowestPrice = prices[0]
        for i in range(1, len(prices)):
            currentPrice = prices[i]
            currentProfit = currentPrice - lowestPrice
            if currentPrice < lowestPrice:
                lowestPrice = currentPrice
            elif currentProfit > bestProfit:
                bestProfit = currentProfit
        return bestProfit
        