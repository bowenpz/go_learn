package daily

func escapeGhosts(ghosts [][]int, target []int) bool {
	cost := distance([]int{0 ,0}, target)
	for _, ghost := range ghosts {
		if cost >= distance(ghost, target) {
			return false
		}
	}
	return true
}

func distance(a []int, b []int) (cost int) {
	if a[0] > b[0] {
		cost += a[0] - b[0]
	} else {
		cost += b[0] - a[0]
	}
	if a[1] > b[1] {
		cost += a[1] - b[1]
	} else {
		cost += b[1] - a[1]
	}
	return
}
