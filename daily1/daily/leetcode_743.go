package daily

type Path struct {
	to, cost int
}

func networkDelayTime(times [][]int, n int, k int) int {
	paths := make(map[int][]*Path)
	for _, time := range times {
		path := &Path{time[1], time[2]}
		if list, ok := paths[time[0]]; ok {
			paths[time[0]] = append(list, path)
		} else {
			paths[time[0]] = []*Path{path}
		}
	}

	costMap := make(map[int]int)
	var dfs func(from int, cost int)
	dfs = func(from int, cost int) {
		if preCost, ok := costMap[from]; !ok || cost < preCost {
			costMap[from] = cost
			for _, path := range paths[from] {
				dfs(path.to, cost+path.cost)
			}
		}
	}
	dfs(k, 0)

	minCost, pathN := -1, 0
	for _, cost := range costMap {
		if minCost < cost {
			minCost = cost
		}
		pathN++
	}

	if pathN == n {
		return minCost
	} else {
		return -1
	}
}
