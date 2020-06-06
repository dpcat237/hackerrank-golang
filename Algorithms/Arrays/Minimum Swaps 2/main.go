package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minSwaps())
}

func minSwaps() uint64 {
	var tt uint64
	if _, err := fmt.Scanf("%d", &tt); err != nil {
		return 0
	}

	line := readLine(bufio.NewReaderSize(os.Stdin, 1024*1024))
	numsStr := strings.Split(line, " ")
	var nums []uint64
	for _, nStr := range numsStr {
		n, err := strconv.ParseInt(nStr, 10, 64)
		if err != nil {
			return 0
		}
		nums = append(nums, uint64(n))
	}

	_, s := minSort(nums, tt)

	return s
}

func findMin(nums []uint64) uint64 {
	var min uint64
	for _, n := range nums {
		if min == 0 || n < min {
			min = n
		}
	}
	return min
}

func minSort(nums []uint64, tt uint64) ([]uint64, uint64) {
	min := findMin(nums)
	var sorted bool
	var s, sNow uint64

	for !sorted {
		sNow = 0
		for i := uint64(0); i < tt; i++ {
			if nums[i] > i+min {
				s++
				sNow++
				nums[nums[i]-min], nums[i] = nums[i], nums[nums[i]-min]
			}
		}

		if sNow == 0 {
			sorted = true
		}
	}

	return nums, s
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
