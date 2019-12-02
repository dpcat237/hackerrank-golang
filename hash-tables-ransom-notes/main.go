package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	yes := "Yes"
	no := "No"
	magazineT := len(magazine)
	noteT := len(note)
	min := 1
	maxWord := 5
	maxWords := 30000
	if magazineT < min || magazineT > maxWords || noteT < min || noteT > maxWords {
		fmt.Println(no)
		return
	}

	magazineM, ok := createMap(magazine)
	if !ok {
		fmt.Println(no)
		return
	}

	for _, word := range note {
		wLength := len(word)
		if wLength < min || wLength > maxWord {
			fmt.Println(no)
			return
		}
		wHash := stringHash(word)
		if c, exists := magazineM[wHash]; exists && c > 0 {
			magazineM[wHash]--
			continue
		}
		fmt.Println(no)
		return
	}
	fmt.Println(yes)
}

func createMap(items []string) (map[uint32]int, bool) {
	min := 1
	max := 5
	rslt := make(map[uint32]int)
	for _, it := range items {
		length := len(it)
		if length < min || length > max {
			return rslt, false
		}
		rslt[stringHash(it)]++
	}
	return rslt, true
}

func stringHash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
