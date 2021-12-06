package daily

//func isCovered(ranges [][]int, left int, right int) bool {
//	sort.Slice(ranges, func(i, j int) bool {
//		if ranges[i][0] != ranges[j][0] {
//			return ranges[i][0] < ranges[j][0]
//		}
//		return ranges[i][1] > ranges[j][1]
//	})
//
//	l, r := ranges[0][0], ranges[0][1]
//	for _, arr := range ranges {
//		if arr[0] <= r+1 {
//			r = max(r, arr[1])
//		} else {
//			l, r = arr[0], arr[1]
//		}
//
//		if l <= left && r >= right {
//			return true
//		} else if l > left {
//			return false
//		}
//	}
//	return false
//}
//
//func max(a, b int) int {
//	if a < b {
//		return b
//	}
//	return a
//}

// 由于 1 <= ranges.length <= 50 并且 1 <= left <= right <= 50
// 所以可以遍历
func isCovered(ranges [][]int, left int, right int) bool {
	fill := new([51]bool)
	for _, arr := range ranges {
		for i := arr[0]; i <= arr[1]; i++ {
			fill[i] = true
		}
	}
	for i := left; i <= right; i++ {
		if !fill[i] {
			return false
		}
	}
	return true
}