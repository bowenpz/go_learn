package daily

import "fmt"

func sortArray(nums []int) []int {
	doSortArray(nums, 0, len(nums)-1)
	return nums
}

func doSortArray(nums []int, l int, r int) {
	if l < r {
		mid := l + (r-l)/2
		doSortArray(nums, l, mid)
		doSortArray(nums, mid+1, r)
		merge(nums, l, r, mid)
	}
}

func merge(nums []int, l int, r int, mid int) {
	help := make([]int, r-l+1, r-l+1)
	copyPart(help, nums, l, r)
	for i, j, k := l, mid+1, l; i <= mid || j <= r; k++ {
		if i > mid {
			nums[k] = help[j-l]
			j++
		} else if j > r {
			nums[k] = help[i-l]
			i++
		} else {
			if help[i-l] < help[j-l] {
				nums[k] = help[i-l]
				i++
			} else {
				nums[k] = help[j-l]
				j++
			}
		}
	}
}

func copyPart(help []int, nums []int, l int, r int) {
	curr := l
	for curr <= r {
		help[curr-l] = nums[curr]
		curr++
	}
}

func main() {
	nums := []int{-2, 3, -5}
	sortArray(nums)
	fmt.Printf("%v\n", nums)
}
