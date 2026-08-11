class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int bestProfit = 0;
        int lowestPrice = prices[0];
        for (int i = 1; i < prices.size(); i++) {
            int currentPrice = prices[i];
            int currentProfit = currentPrice - lowestPrice;
            if (currentPrice < lowestPrice) {
                lowestPrice = currentPrice;
            } else if (currentProfit > bestProfit) {
                bestProfit = currentProfit;
            }
        }
        return bestProfit;
    }
};
