package daily

func allPathsSourceTarget(graph [][]int) (ans [][]int) {
	var dfs func(start, end int, path []int)
	dfs = func(start, end int, path []int) {
		path = append(path, start)
		if start == end {
			ans = append(ans, append([]int(nil), path...))
		} else {
			for _, next := range graph[start] {
				dfs(next, end, path)
			}
		}
		path = path[:len(path)-1]
	}
	dfs(0, len(graph)-1, []int{})
	return
}
