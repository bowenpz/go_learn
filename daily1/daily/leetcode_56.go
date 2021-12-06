package daily

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[j][1] < intervals[i][1]
		}
	})
	l, r := intervals[0][0], intervals[0][1]
	ans := make([][]int, 0)
	for _, interval := range intervals[1:] {
		if interval[0] <= r {
			if r < interval[1] {
				r = interval[1]
			}
		} else {
			ans = append(ans, []int{l, r})
			l, r = interval[0], interval[1]
		}
	}
	return append(ans, []int{l, r})
}