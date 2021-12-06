package daily

import "math"

func findUnsortedSubarray(nums []int) int {
	const UnReach = math.MaxInt16
	peak, valley := UnReach, UnReach
	for i, num := range nums {
		if i > 0 && num < nums[i-1] {
			if valley == UnReach || valley > num {
				valley = num
			}
			if peak == UnReach || peak < nums[i-1] {
				peak = nums[i-1]
			}
		}
	}

	if peak == UnReach {
		return 0
	}

	l, r := 0, len(nums)-1
	for l < len(nums) {
		if nums[l] > valley {
			break
		}
		l++
	}
	for r >= 0 {
		if nums[r] < peak {
			break
		}
		r--
	}
	return r - l + 1
}
