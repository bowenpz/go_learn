package main

import (
	"math/rand"
	"time"
)

type Solution struct {
	interval int
	sum      []int
}

func Constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	sum := make([]int, len(w))
	for i, num := range w {
		if i == 0 {
			sum[i] = num
		} else {
			sum[i] = num + sum[i-1]
		}
	}
	return Solution{sum: sum, interval: sum[len(sum)-1]}
}

func (s *Solution) PickIndex() int {
	randNum := rand.Intn(s.interval) + 1
	l, r := 0, len(s.sum)-1
	for l < r {
		mid := l + (r-l)/2
		if randNum == s.sum[mid] {
			l = mid
			break
		} else if randNum > s.sum[mid] {
			l = mid + 1
		} else {
			if mid == 0 || randNum > s.sum[mid-1] {
				l = mid
				break
			} else {
				r = mid - 1
			}
		}
	}
	return s.sum[l]
}
