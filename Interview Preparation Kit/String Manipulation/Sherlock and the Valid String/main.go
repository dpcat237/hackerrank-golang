package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func isValidSherlock(a string) string {
	if len(a) == 0 {
		return "NO"
	}
	if len(a) == 1 {
		return "YES"
	}

	aInt := mergeSort(toNumArray(a))
	rst := countValidSherlock(aInt, 0, 0, 0, 0, false)

	if rst {
		return "YES"
	}
	return "NO"
}

func countValidSherlock(a []uint64, freq, cFreq, prev, cur uint64, removed bool) bool {
	if len(a) == 0 {
		return freq == cFreq || (!removed && (freq == cFreq-1 || freq == cFreq+1 || cFreq == 1))
	}

	if prev == 0 && a[0] != cur {
		return countValidSherlock(a[1:], cFreq, 1, cur, a[0], false)
	}

	if a[0] == cur {
		if freq == 0 || freq > cFreq {
			return countValidSherlock(a[1:], freq, cFreq+1, prev, cur, removed)
		}

		// when freq == cFreq
		if removed {
			return false
		}
		return countValidSherlock(a[1:], freq, cFreq+1, prev, cur, true)
	}

	// when a[0] != cur
	if freq > cFreq+1 {
		return false
	}
	if freq > cFreq {
		return countValidSherlock(a[1:], freq, 1, cur, a[0], true)
	}
	return countValidSherlock(a[1:], freq, 1, cur, a[0], removed)
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

	s := readLine(reader)

	result := isValidSherlock(s)

	fmt.Fprintf(writer, "%s\n", result)

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
