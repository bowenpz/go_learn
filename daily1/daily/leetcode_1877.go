package daily

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	max := 0
	for i := 0; i < n; i++ {
		if max < nums[i]+nums[n-i-1] {
			max = nums[i] + nums[n-i-1]
		}
	}
	return max
}
