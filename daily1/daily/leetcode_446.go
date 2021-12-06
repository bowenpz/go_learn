package daily

func numberOfArithmeticSlices2(nums []int) (ans int) {
	dp := make([]map[int]int, len(nums))
	for j, y := range nums {
		dp[j] = map[int]int{}
		for i, x := range nums[:j] {
			d := y - x
			ans += dp[i][d]
			dp[j][d] += dp[i][d] + 1
		}
	}
	return
}