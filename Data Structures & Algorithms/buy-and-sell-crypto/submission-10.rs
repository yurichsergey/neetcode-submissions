impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.is_empty() {
            return 0;
        }
        let mut best_profit = 0;
        let mut lowest_price = prices[0];
        for i in 1..prices.len() {
            let current_price = prices[i];
            let current_profit = current_price - lowest_price;
            lowest_price = lowest_price.min(current_price);
            best_profit = best_profit.max(current_profit);
        }
        best_profit
    }
}
