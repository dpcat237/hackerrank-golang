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
	fmt.Println(rectangle())
}

func rectangle() int {
	var tt int
	if _, err := fmt.Scanf("%d", &tt); err != nil {
		return 0
	}

	line := readLine(bufio.NewReaderSize(os.Stdin, 1024*1024))
	numsStr := strings.Split(line, " ")
	var max int
	for i := 1; i < tt-1; i++ {
		n, err := strconv.ParseInt(numsStr[i], 10, 64)
		if err != nil {
			return 0
		}

		lst := largest(i, int(n), tt, numsStr)
		rct := int(n) * lst
		if rct > max {
			max = rct
		}
	}

	return max
}

func largest(k, n, tt int, numsStr []string) int {
	c := 1
	for i := k - 1; i > 0; i-- {
		ln, err := strconv.ParseInt(numsStr[i], 10, 64)
		if err != nil {
			return 0
		}
		if int(ln) >= n {
			c++
			continue
		}
		break
	}

	for i := k + 1; i < tt; i++ {
		rn, err := strconv.ParseInt(numsStr[i], 10, 64)
		if err != nil {
			return 0
		}
		if int(rn) >= n {
			c++
			continue
		}
		break
	}

	return c
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
