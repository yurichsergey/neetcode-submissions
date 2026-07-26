class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int bestProfit = 0;
        int lowestPrice = prices[0];
        for (int i = 1; i < prices.size(); i++) {
            int currentPrice = prices[i];
            int currentProfit = currentPrice - lowestPrice;

            lowestPrice = min(lowestPrice, currentPrice);
            bestProfit = max(bestProfit, currentProfit);
        }
        return bestProfit;
    }
};
