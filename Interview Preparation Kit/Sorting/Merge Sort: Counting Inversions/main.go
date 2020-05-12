package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	countInversions()
}

func countInversions() {
	in := bufio.NewReaderSize(os.Stdin, 1024*1024)
	for {
		inputBt, _, err := in.ReadLine()
		input := string(inputBt)
		if err != nil {
			break
		}

		processLine(input)
	}
}

func processLine(input string) {
	numsStr := strings.Split(strings.TrimSpace(input), " ")
	if len(numsStr) < 2 {
		return
	}

	_, inv := mergeSortInversions(toNumArray(numsStr))
	if _, err := fmt.Fprintf(os.Stdout, "%d\n", inv); err != nil {
		return
	}
}

func toNumArray(numsStr []string) []uint64 {
	var nums []uint64
	for _, numStr := range numsStr {
		n, err := strconv.Atoi(numStr)
		if err != nil {
			return []uint64{}
		}
		nums = append(nums, uint64(n))
	}
	return nums
}

func mergeSortInversions(arr []uint64) ([]uint64, uint64) {
	if len(arr) == 1 {
		return arr, 0
	}

	l := arr[:len(arr)/2]
	r := arr[len(arr)/2:]

	l, ai := mergeSortInversions(l)
	r, bi := mergeSortInversions(r)

	var rst []uint64
	var i, j uint64
	inv := 0 + ai + bi

	for i < uint64(len(l)) && j < uint64(len(r)) {
		if l[i] <= r[j] {
			rst = append(rst, l[i])
			i++
			continue
		}
		rst = append(rst, r[j])
		inv += uint64(len(l)) - i
		j++
	}

	rst = append(rst, l[i:]...)
	rst = append(rst, r[j:]...)

	return rst, inv
}
