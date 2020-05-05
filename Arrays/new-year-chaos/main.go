package main

import (
	"fmt"
)

func main() {
	minimumBribes([]int32{1, 2, 5, 3, 7, 8, 6, 4})
}

func minimumBribes(q []int32) {
	count := int32(0)
	for i := int32(len(q)) - 1; i >= 0; i-- {
		if q[i]-(i+1) > 2 {
			fmt.Println("Too chaotic")
			return
		}
		for j := max(0, q[i]-2); j < i; j++ {
			if q[j] > q[i] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func max(a, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}
