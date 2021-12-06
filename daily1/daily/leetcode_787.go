package daily

// FindCheapestPrice error
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cost := make([][]int, n)
	went := make([][]bool, n)
	for i := 0; i < n; i++ {
		cost[i] = make([]int, n)
		went[i] = make([]bool, n)
	}

	next := make([][]int, n)
	for _, flight := range flights {
		next[flight[0]] = append(next[flight[0]], flight[1])
		cost[flight[0]][flight[1]] = flight[2]
	}

	var dfs func(src, dst, k int) int
	dfs = func(src, dst, k int) (minCost int) {
		if src == dst || k < 0 {
			return
		}
		minCost = cost[src][dst]
		for _, nextFlight := range next[src] {
			if !went[src][nextFlight] {
				went[src][nextFlight] = true

				if nextFlight != src {
					nextCost := dfs(nextFlight, dst, k-1)
					if currCost := cost[src][nextFlight] + nextCost; nextCost > 0 && currCost < minCost {
						minCost = currCost
					}
				}
			}
		}
		cost[src][dst] = minCost
		return
	}
	dfs(src, dst, k)

	if cost[src][dst] == 0 {
		return -1
	} else {
		return cost[src][dst]
	}
}
