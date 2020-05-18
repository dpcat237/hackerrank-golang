package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func makeAnagram(a string, b string) int32 {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if (len(a) == 1 && len(b) == 1) && a[0] == b[0] {
		return 0
	}

	aInt := mergeSort(toNumArray(a))
	bInt := mergeSort(toNumArray(b))

	return int32(countAnagram(aInt, bInt, 0))
}

func countAnagram(a, b []uint64, c uint64) uint64 {
	if len(a) == 0 {
		return c + uint64(len(b))
	}
	if len(b) == 0 {
		return c + uint64(len(a))
	}
	if (len(a) == 1 && len(b) == 1) && a[0] == b[0] {
		return c
	}

	if a[0] == b[0] {
		return countAnagram(a[1:], b[1:], c)
	}
	if a[0] > b[0] {
		return countAnagram(a, b[1:], c+1)
	}
	return countAnagram(a[1:], b, c+1)
}

func toNumArray(numsStr string) []uint64 {
	var nums []uint64
	for _, numStr := range numsStr {
		nums = append(nums, uint64(numStr))
	}
	return nums
}

func mergeSort(arr []uint64) []uint64 {
	if len(arr) == 1 {
		return arr
	}

	l := arr[:len(arr)/2]
	r := arr[len(arr)/2:]

	l = mergeSort(l)
	r = mergeSort(r)

	var rst []uint64
	var i, j uint64

	for i < uint64(len(l)) && j < uint64(len(r)) {
		if l[i] <= r[j] {
			rst = append(rst, l[i])
			i++
			continue
		}
		rst = append(rst, r[j])
		j++
	}

	rst = append(rst, l[i:]...)
	rst = append(rst, r[j:]...)

	return rst
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
