impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.is_empty() {
            return 0;
        }
        let mut bestProfit = 0;
        let mut lowestPrice = prices[0];
        for i in 1..prices.len() {
            let currentPrice = prices[i];
            let currentProfit = currentPrice - lowestPrice;
            lowestPrice = lowestPrice.min(currentPrice);
            bestProfit = bestProfit.max(currentProfit);
        }
        bestProfit
    }
}
