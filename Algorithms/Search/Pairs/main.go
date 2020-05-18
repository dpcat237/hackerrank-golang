package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func pairs(k int64, arr []int64) int64 {
	var rst int64
	arrMp := arrayToMap(arr)
	for _, n := range arr {
		diff := n - k

		if _, ok := arrMp[diff]; ok {
			rst++
		}
	}
	return rst
}

func arrayToMap(arr []int64) map[int64]bool {
	mp := make(map[int64]bool)
	for _, a := range arr {
		if _, ok := mp[a]; !ok {
			mp[a] = true
		}
	}
	return mp
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := kTemp

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := arrItemTemp
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

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
