package daily


func pathInZigZagTree(label int) []int {
	path := make([]int, 0)

	for label > 0 {
		level, i := getIndex(label)
		path = append(path, label)
		label = getLabel(level-1, (i+1)/2)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// 数字 -> 第几层第几个
func getIndex(label int) (level, i int) {
	curr := label
	for curr > 0 {
		curr >>= 1
		level++
	}

	// 正向：true  反向：false
	if flag := level%2 == 1; flag {
		start := 1 << (level - 1)
		i = label - start + 1
	} else {
		start := 1 << level
		i = start - label
	}

	return
}

// 第几层第几个 -> 数字
func getLabel(level, i int) int {
	if level == 0 {
		return 0
	}

	// 该层最小值
	start := 1 << (level - 1)

	// 正向：true  反向：false
	if flag := level%2 == 1; flag {
		return start + i - 1
	} else {
		return (start << 1) - i
	}
}
