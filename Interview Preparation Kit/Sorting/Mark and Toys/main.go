package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sortToys()
}

type Data struct {
	canBuy     int32
	totalToys  int32
	totalFunds int32
	toys       []int
}

func sortToys() {
	var d Data

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	d.setTotals(readLine(reader))
	if err := d.setToys(readLine(reader)); err != nil {
		d.printResult()
		return
	}

	if d.totalFunds == 0 || len(d.toys) == 0 {
		d.printResult()
	}

	d.countToBuy()
	d.printResult()
}

func (d *Data) countToBuy() {
	var sum int32
	for _, toy := range d.toys {
		sum += int32(toy)
		if sum <= d.totalFunds {
			d.canBuy++
		}

		if sum >= d.totalFunds {
			return
		}
	}
}

func (d Data) printResult() {
	_, err := fmt.Fprintf(os.Stdout, "%d", d.canBuy)
	if err != nil {
		return
	}
}

func (d *Data) setTotals(input string) {
	data := strings.Split(input, " ")
	tTotal, err := strconv.ParseInt(data[0], 10, 32)
	if err != nil {
		return
	}
	fTotal, err := strconv.ParseInt(data[1], 10, 32)
	if err != nil {
		return
	}

	d.totalToys = int32(tTotal)
	d.totalFunds = int32(fTotal)
}

func (d *Data) setToys(input string) error {
	for _, toy := range strings.Split(input, " ") {
		toyP, err := strconv.ParseInt(toy, 10, 64)
		if err != nil {
			return err
		}
		d.toys = append(d.toys, int(toyP))
	}
	sort.Ints(d.toys)
	return nil
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
