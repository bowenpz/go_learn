package daily

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))

	chars := []byte(s)
	for i := len(chars) - 1; i >= 0; i-- {
		dp[i] = make([]int, len(chars))
		dp[i][i] = 1
		for j := i + 1; j < len(chars); j++ {
			if chars[i] == chars[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(chars)-1]
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}