class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices: number[]): number {
        let bestProfit = 0;
        let lowestPrice = prices[0];
        for (let i = 0; i < prices.length; i++) {
            const currentPrice = prices[i];
            const currentProfit = currentPrice - lowestPrice;
            if (currentPrice < lowestPrice) {
                lowestPrice = currentPrice;
            } else if (currentProfit > bestProfit) {
                bestProfit = currentProfit;
            }
        }
        return bestProfit;
    }
}
