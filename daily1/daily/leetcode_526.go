package daily

func countArrangement(n int) int {
	var dfs func([]int, []bool, int) int
	dfs = func(arr []int, used []bool, i int) int {
		if i > n {
			return 1
		}
		count := 0
		for j := 1; j <= n; j++ {
			if !used[j] && divided(i, j) {
				used[j] = true
				arr[i] = j
				count += dfs(arr, used, i + 1)
				used[j] = false
			}
		}
		return count
	}
	return dfs(make([]int, n+1), make([]bool, n+1), 1)
}

func divided(i int, j int) bool {
	if i > j {
		return i%j == 0
	}
	return j%i == 0
}
