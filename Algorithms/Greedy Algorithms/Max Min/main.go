package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMax())
}

func minMax() int {
	var tt, k int
	if _, err := fmt.Scanf("%d", &tt); err != nil {
		return 0
	}
	if _, err := fmt.Scanf("%d", &k); err != nil {
		return 0
	}

	var nums []int
	for {
		var n int
		if _, err := fmt.Scanf("%d", &n); err != nil {
			break
		}
		nums = append(nums, n)
	}
	if len(nums) < k {
		return 0
	}

	sort.Ints(nums)
	return minimumUnfairness(nums, k)
}

func minimumUnfairness(nums []int, k int) int {
	tt := len(nums)
	unf := nums[tt-1]

	for i := 0; i <= tt-k; i++ {
		mm := nums[i+k-1] - nums[i]
		if mm < unf {
			unf = mm
		}
	}
	return unf
}
