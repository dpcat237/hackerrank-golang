package main

import "fmt"

func main() {
	fmt.Println(arrayMaximum())
}

func arrayMaximum() uint64 {
	var sz, opTt uint64
	if _, err := fmt.Scanf("%d %d", &sz, &opTt); err != nil {
		return 0
	}

	arr := make([]uint64, sz+1)
	for {
		var k1, k2, val uint64
		if _, err := fmt.Scanf("%d %d %d", &k1, &k2, &val); err != nil {
			break
		}

		k1--
		arr[k1] += val
		arr[k2] -= val
	}

	var val, mx uint64
	for k := range arr {
		val += arr[k]
		if val > mx {
			mx = val
		}
	}

	return mx
}
