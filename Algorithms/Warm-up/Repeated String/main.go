package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	min := 1
	maxS := 100
	maxN := int64(1000000000000)
	sTotal := len(s)
	if sTotal < min || sTotal > maxS || n < 1 || n > maxN {
		return 0
	}

	asCount := countA(sTotal, s)
	if asCount == 0 {
		return 0
	}
	if asCount == sTotal {
		return n
	}

	rst := (n / int64(sTotal)) * int64(asCount)
	modT := n % int64(sTotal)
	if modT > 0 {
		rst += int64(countA(int(modT), s[:modT]))
	}
	return rst
}

func countA(sTotal int, s string) int {
	return sTotal - len(strings.Replace(s, "a", "", -1))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
