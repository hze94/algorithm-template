/* JavaScript */
// LeetCode188 买卖股票的最佳时机IV
/**
 * @param {number} c
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (c, prices) {
  let n = prices.length;
  // 0. Move index to 1-based
  prices.unshift(0);
  // 1. Define f, initiailize -oo
  // 数值类型可以直接fill，嵌套的数组类型要使用map，不然n+1个数组会指向同一个对象
  let f = Array(n + 1).fill().map(
    () => Array(2).fill().map(
      () => Array(c + 1).fill(-1e9)));
  f[0][0][0] = 0;
  // 2. Loop over all states
  for (let i = 1; i <= n; i++)
    for (let j = 0; j < 2; j++)
      for (let k = 0; k <= c; k++) {
        // 3. Copy decisions
        if (k > 0) f[i][1][k] = Math.max(f[i][1][k], f[i - 1][0][k - 1] - prices[i]);
        f[i][0][k] = Math.max(f[i][0][k], f[i - 1][1][k] + prices[i]);
        f[i][j][k] = Math.max(f[i][j][k], f[i - 1][j][k]);
      }
  // 4. return target
  let ans = -1e9;
  for (let k = 0; k <= c; k++)
    ans = Math.max(ans, f[n][0][k]);
  return ans;
};
