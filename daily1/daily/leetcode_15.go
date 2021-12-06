package daily

import "sort"

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for i, num := range nums {
		if i == 0 || num != nums[i-1] {
			if num > 0 {
				break
			}
			for l, r := i+1, len(nums)-1; l < r; {
				if l > i+1 && nums[l] == nums[l-1] {
					l++
				} else {
					sum := num + nums[l] + nums[r]
					if sum == 0 {
						ans = append(ans, []int{num, nums[l], nums[r]})
						l++
					} else if sum > 0 {
						r--
					} else {
						l++
					}
				}
			}
		}
	}
	return
}
