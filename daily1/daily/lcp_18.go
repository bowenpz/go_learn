package daily

import "sort"

const MOD = 1e9 + 7

func breakfastNumber(staple []int, drinks []int, x int) (sum int) {
	sort.Ints(staple)
	sort.Ints(drinks)
	for i, j := 0, len(drinks)-1; i < len(staple) && j >= 0; i++ {
		for j >= 0 && staple[i]+drinks[j] > x {
			j--
		}
		if j >= 0 {
			sum += j + 1
			sum %= MOD
		}
	}
	return
}
