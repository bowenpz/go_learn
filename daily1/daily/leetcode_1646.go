package daily

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	nums := make([]int, n+1)
	nums[0], nums[1] = 0, 1
	max := 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[i/2+1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
