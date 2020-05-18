package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	yes := "YES"
	no := "NO"
	min := 1
	max := 100000
	s1T := len(s1)
	s2T := len(s2)
	if s1T < min || s1T > max || s2T < min || s2T > max {
		return no
	}

	s1M, ok := stringToMap(s1)
	if !ok {
		return no
	}

	for _, s := range s2 {
		if _, exists := s1M[int(s)]; exists {
			return yes
		}
	}

	return no
}

func stringToMap(str string) (map[int]bool, bool) {
	min := 97
	max := 122
	rst := make(map[int]bool)
	for _, s := range str {
		sInt := int(s)
		if sInt < min || sInt > max {
			return rst, false
		}
		rst[sInt] = true
	}
	return rst, true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
