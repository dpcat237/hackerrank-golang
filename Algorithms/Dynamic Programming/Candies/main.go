package main

import "fmt"

func main() {
	candies()
}

func candies() {
	var tt int
	if _, err := fmt.Scanf("%d", &tt); err != nil {
		return
	}

	rts := make([]int, tt)
	for i := 0; i < tt; i++ {
		if _, err := fmt.Scanf("%d", &rts[i]); err != nil {
			return
		}
	}
	fmt.Println(minRequired(rts))
}

func minRequired(rts []int) int {
	tt := len(rts)
	minR := make([]int, tt)
	minL := make([]int, tt)

	for i := 0; i < tt; i++ {
		minR[i] = 1
		minL[i] = 1
	}

	for i := tt - 2; i >= 0; i-- {
		if rts[i] > rts[i+1] {
			minR[i] = 1 + minR[i+1]
		}
	}

	for i := 1; i < tt; i++ {
		if rts[i-1] < rts[i] {
			minL[i] = 1 + minL[i-1]
		}
	}

	var c int
	for i := 0; i < tt; i++ {
		if minL[i] < minR[i] {
			c = c + minR[i]
			continue
		}
		c = c + minL[i]
	}
	return c
}
