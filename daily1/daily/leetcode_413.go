package daily

func numberOfArithmeticSlices(nums []int) int {
	distances := make([]int, len(nums))
	distances[0] = 2001

	continuous, sum := 0, 0
	for i := 1; i < len(nums); i++ {
		distances[i] = nums[i] - nums[i-1]
		if distances[i] == distances[i-1] {
			continuous++
			sum += continuous
		} else {
			continuous = 0
		}
	}

	return sum
}