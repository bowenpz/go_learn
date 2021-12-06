package daily

func checkRecord1(n int) int {
	// dp[i][j][k]：到第 i 天为止，发生了 j 次缺勤，连续发生了 k 次迟到
	dp := make([][2][3]int, n+1)
	dp[1][0][0], dp[1][1][0], dp[1][0][1] = 1, 1, 1

	for i := 2; i <= n; i++ {
		for k := 0; k < 3; k++ {
			if k == 0 { // 连续 0 次迟到：这次不是迟到，可能是缺勤，也可能是到场
				// 0 次缺勤
				dp[i][0][k] = add(dp[i-1][0][0], dp[i-1][0][1], dp[i-1][0][2])
				// 1 次缺勤
				dp[i][1][k] = add(dp[i-1][1][0], dp[i-1][1][1], dp[i-1][1][2],
					dp[i-1][0][0], dp[i-1][0][1], dp[i-1][0][2])
			} else { // 连续 k 次迟到：这次就是迟到
				// 0 次缺勤
				dp[i][0][k] = dp[i-1][0][k-1]
				// 1 次缺勤
				dp[i][1][k] = dp[i-1][1][k-1]
			}
		}
	}

	ans := 0
	for _, arr := range dp[n] {
		for _, num := range arr {
			ans = add(ans, num)
		}
	}
	return ans
}

func add(nums ...int) (sum int) {
	const mod = 1e9 + 7
	for _, num := range nums {
		sum += num
		sum %= mod
	}
	return
}
