package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(findNearest())
}

func findNearest() int64 {
	rst := int64(-1)
	bufin := bufio.NewReader(os.Stdin)

	var q, eg int
	if _, err := fmt.Sscanf(readLine(bufin), "%d%d", &q, &eg); err != nil {
		return rst
	}

	if eg <= 1 {
		return rst
	}

	for i := 0; i < eg; i++ {
		readLine(bufin)
	}

	crs := readLine(bufin)
	var fCr string
	if _, err := fmt.Sscanf(readLine(bufin), "%s", &fCr); err != nil {
		return rst
	}

	var dist, minDist int64
	for _, cl := range strings.Split(crs, " ") {
		if cl == fCr {
			if dist == 0 {
				dist++
				continue
			}

			if dist == 1 {
				minDist = dist
				break
			}

			if minDist > 0 {
				minDist = min(minDist, dist)
				dist = 1
				continue
			}

			minDist = dist
			dist = 1
			continue
		}

		if dist != 0 {
			dist++
		}
	}

	if minDist == 0 {
		return rst
	}

	return minDist
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
