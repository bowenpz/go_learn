package daily

func minimumOperations(leaves string) int {
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(leaves)
	dp := make([][3]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			if leaves[i] == 'r' {
				dp[i][0] = 0
			} else {
				dp[i][0] = 1
			}
		} else {
			if leaves[i] == 'r' {
				dp[i][0] = dp[i-1][0]
			} else {
				dp[i][0] = dp[i-1][0] + 1
			}
		}

		if i == 1 {
			if leaves[i] == 'r' {
				dp[i][1] = dp[i-1][0] + 1
			} else {
				dp[i][1] = dp[i-1][0]
			}
		} else if i > 1 {
			if leaves[i] == 'r' {
				dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
			} else {
				dp[i][1] = min(dp[i-1][0], dp[i-1][1])
			}
		}

		if i == 2 {
			if leaves[i] == 'r' {
				dp[i][2] = dp[i-1][1]
			} else {
				dp[i][2] = dp[i-1][1] + 1
			}
		} else if i > 2 {
			if leaves[i] == 'r' {
				dp[i][2] = min(dp[i-1][1], dp[i-1][2])
			} else {
				dp[i][2] = min(dp[i-1][1], dp[i-1][2]) + 1
			}
		}
	}

	return dp[n-1][2]
}