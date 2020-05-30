package main

import "fmt"

func main() {
	roadsAndLibraries()
}

func roadsAndLibraries() {
	var q int
	if _, err := fmt.Scanf("%d", &q); err != nil {
		return
	}

	for line := 0; line < q; line++ {
		var n, m int
		var pRoad, pLib int64
		if _, err := fmt.Scanf("%d%d%d%d", &n, &m, &pLib, &pRoad); err != nil {
			return
		}
		processProvince(n, m, pLib, pRoad)
	}
}

func find(p []int, v int) int {
	if p[v] == v {
		return v
	}
	p[v] = find(p, p[v])
	return p[v]
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func processProvince(n, m int, pLib, pRoad int64) {
	var p []int
	for i := 0; i < n; i++ {
		p = append(p, i)
	}
	var cRoad int64
	for i := 0; i < m; i++ {
		u, v := 0, 0

		if _, err := fmt.Scanf("%d%d", &u, &v); err != nil {
			return
		}

		u, v = u-1, v-1
		if (i & 1) == 0 {
			u, v = v, u
		}
		u, v = find(p, u), find(p, v)
		if u != v {
			p[u] = v
			cRoad++
		}
	}
	fmt.Printf("%d\n", min(int64(n)*pLib, cRoad*pRoad+(int64(n)-cRoad)*pLib))
}
