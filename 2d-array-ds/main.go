package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	hTotal := len(arr)
	if hTotal < 3 || hTotal > 6 {
		return 0
	}

	rst := int32(-100)
	for i := 1; i < hTotal-1; i++ {
		sum, ch := countSum(arr[i-1], arr[i], arr[i+1])
		if !ch {
			return 0
		}
		if sum > rst {
			rst = sum
		}
	}
	return rst
}

func countSum(aRr, bRr, cRr []int32) (int32, bool) {
	rst := int32(-100)
	for i := 1; i < len(aRr)-1; i++ {
		a1 := aRr[i-1]
		a2 := aRr[i]
		a3 := aRr[i+1]
		b2 := bRr[i]
		c1 := cRr[i-1]
		c2 := cRr[i]
		c3 := cRr[i+1]
		if !validValue(a1) || !validValue(a2) || !validValue(a3) || !validValue(b2) || !validValue(c1) || !validValue(c2) || !validValue(c3) {
			return 0, false
		}
		sum := a1 + a2 + a3 + b2 + c1 + c2 + c3
		if sum > rst {
			rst = sum
		}
	}
	return rst, true
}

func validValue(n int32) bool {
	min := -9
	max := 9
	return int(n) >= min && int(n) <= max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
