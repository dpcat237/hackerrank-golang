package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Query struct {
	ByNum  map[int32]int32
	ByFreq map[int32]*Freq
	Result []int32
}

type Freq struct {
	Nums map[int32]bool
}

func freqQuery(queries [][]int32) []int32 {
	var fq Query
	fq.init()
	for _, q := range queries {
		if len(q) == 1 {
			continue
		}
		fq.process(q[0], q[1])
	}
	return fq.Result
}

func (fq *Query) init() {
	fq.ByNum = make(map[int32]int32)
	fq.ByFreq = make(map[int32]*Freq)
}

func (fq *Query) process(q, num int32) {
	switch q {
	case 1:
		fq.insert(num)
	case 2:
		fq.remove(num)
	case 3:
		fq.query(num)
	}
}

func (fq *Query) insert(num int32) {
	if _, ok := fq.ByNum[num]; !ok {
		fq.ByNum[num] = 1
		fq.insertFreq(num, 1)
		return
	}

	freq := fq.ByNum[num]
	(*fq.ByFreq[freq]).remove(num)
	if (*fq.ByFreq[freq]).count() == 0 {
		var f Freq
		f.init()
		fq.ByFreq[freq] = &f
	}

	fq.ByNum[num]++
	fq.insertFreq(num, fq.ByNum[num])
}

func (fq *Query) insertFreq(num, freq int32) {
	if _, ok := fq.ByFreq[freq]; !ok {
		var f Freq
		f.init()
		f.insert(num)
		fq.ByFreq[freq] = &f
	}
	(*fq.ByFreq[freq]).insert(num)
}

func (f *Freq) init() {
	f.Nums = make(map[int32]bool)
}

func (f *Freq) insert(num int32) {
	if _, ok := f.Nums[num]; !ok {
		f.Nums[num] = true
	}
}

func (f Freq) count() int32 {
	return int32(len(f.Nums))
}

func (fq *Query) query(freq int32) {
	if _, ok := fq.ByFreq[freq]; ok && (*fq.ByFreq[freq]).count() > 0 {
		fq.Result = append(fq.Result, 1)
		return
	}
	fq.Result = append(fq.Result, 0)
}

func (f *Freq) remove(num int32) {
	delete(f.Nums, num)
}

func (fq *Query) remove(num int32) {
	if _, ok := fq.ByNum[num]; !ok {
		return
	}

	freq := fq.ByNum[num]
	(*fq.ByFreq[freq]).remove(num)
	if (*fq.ByFreq[freq]).count() == 0 {
		var f Freq
		f.init()
		fq.ByFreq[freq] = &f
	}

	fq.ByNum[num]--
	if fq.ByNum[num] == 0 {
		delete(fq.ByNum, num)
		return
	}
	fq.insertFreq(num, fq.ByNum[num])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
