// C/C++
// LeetCode188 买卖股票的最佳时机IV
class Solution {
public:
  int maxProfit(int c, vector<int>& prices) {
    int n = prices.size();
    // 0. Move index to 1-based
    prices.insert(prices.begin(), 0);
    // 1. Define f, initiailize -oo
    vector<vector<vector<int>>> f(n + 1, vector<vector<int>>(2, vector<int>(c + 1, -1e9)));
    f[0][0][0] = 0;
    // 2. Loop over all states
    for (int i = 1; i <= n; i++)
      for (int j = 0; j < 2; j++)
        for (int k = 0; k <= c; k++) {
          // 3. Copy decisions
          if (k > 0) f[i][1][k] = max(f[i][1][k], f[i - 1][0][k - 1] - prices[i]);
          f[i][0][k] = max(f[i][0][k], f[i - 1][1][k] + prices[i]);
          f[i][j][k] = max(f[i][j][k], f[i - 1][j][k]);
        }
    // 4. return target
    int ans = -1e9;
    for (int k = 0; k <= c; k++)
      ans = max(ans, f[n][0][k]);
    return ans;
  }
};
