// Go
// LeetCode188 买卖股票的最佳时机IV
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
	
func maxProfit(c int, prices []int) int {
	n := len(prices);
	// 0. Move index to 1-based
	prices = append([]int{0}, prices...)
	
	// 1. Define f, initiailize -oo
	f := make([][][]int, n + 1)
	for i := 0; i <= n; i++ {
		f[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			f[i][j] = make([]int, c + 1)
			for k := 0; k <= c; k++ {
				f[i][j][k] = -1000000000
			}
		}
	}
	f[0][0][0] = 0
	
	// 2. Loop over all states
	for i := 1; i <= n; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k <= c; k++ {
				// 3. Copy decisions
				if k > 0 {
					f[i][1][k] = max(f[i][1][k], f[i - 1][0][k - 1] - prices[i])
				}
				f[i][0][k] = max(f[i][0][k], f[i - 1][1][k] + prices[i])
				f[i][j][k] = max(f[i][j][k], f[i - 1][j][k])
			}
		}
	}
	
	// 4. return target
	ans := -1000000000
	for k := 0; k <= c; k++ {
		ans = max(ans, f[n][0][k])
	}
	return ans
}