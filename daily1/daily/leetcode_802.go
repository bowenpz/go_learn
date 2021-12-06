package daily

type State byte

const (
	Init   State = 0
	Safe   State = 1
	Unsafe State = 2
)

func eventualSafeNodes(graph [][]int) []int {
	states := make([]State, len(graph))

	var dfs func(node int) bool
	dfs = func(node int) bool {
		if states[node] == Init {
			states[node] = Unsafe
			for _, nextNode := range graph[node] {
				if ok := dfs(nextNode); !ok {
					return false
				}
			}
			states[node] = Safe
			return true
		}
		return states[node] == Safe
	}

	ans := make([]int, 0)
	for i := 0; i < len(graph); i++ {
		if ok := dfs(i); ok {
			ans = append(ans, i)
		}
	}
	return ans
}
