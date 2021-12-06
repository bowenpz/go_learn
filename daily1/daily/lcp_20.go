package daily

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	MOD := int(1e9+7)
	memo := make(map[int]int)

	var dfs func(target int) int
	dfs = func(target int) int {
		if target == 1 {
			return inc
		}
		if target == 0 {
			return 0
		}

		if res, ok := memo[target]; ok {
			return res
		}

		res := inc * target
		for i, scale := range jump {
			t, k := target/scale, target%scale
			res = min(res, dfs(t)+cost[i]+k*inc)
			if k > 0 {
				res = min(res, dfs(t+1)+cost[i]+(scale-k)*dec)
			}
		}

		memo[target] = res
		return res
	}
	return dfs(target) % MOD
}
