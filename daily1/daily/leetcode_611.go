package daily

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0

	length := len(nums)
	for i := 0; i < length; i++ {
		for j, k := i+1, i+1; j < length; j++ {
			max := nums[i] + nums[j]
			if k < j {
				k = j
			}
			for k+1 < length {
				if nums[k+1] >= max {
					break
				}
				k++
			}
			count += k - j
		}
	}

	return count
}
