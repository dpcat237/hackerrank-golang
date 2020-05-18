package main

import (
	"fmt"
)

func main() {
	countSwaps([]int32{3, 2, 1})
}

func countSwaps(a []int32) {
	rst, swCount := swap(a, 0)
	fmt.Printf("Array is sorted in %d swaps.\n", swCount)
	fmt.Printf("First Element: %d\n", rst[0])
	fmt.Printf("Last Element: %d\n", rst[len(rst)-1])
}

func swap(a []int32, swCount int32) ([]int32, int32) {
	total := int32(len(a))
	swC := int32(0)
	for i := int32(0); i < total; i++ {
		iNext := i + 1
		if iNext < total && a[i] > a[iNext] {
			sw := a[i]
			a[i] = a[iNext]
			a[iNext] = sw
			swC++
		}
	}

	if swC == 0 {
		return a, swCount
	}
	return swap(a, swCount+swC)
}
