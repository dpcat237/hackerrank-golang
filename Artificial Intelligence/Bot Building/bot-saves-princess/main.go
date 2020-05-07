package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Person struct {
	line     int
	position int
	found    bool
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var m, p Person

	line := 0
	for {
		inputBt, _, err := in.ReadLine()
		input := string(inputBt)
		if err != nil {
			break
		}

		if mi := strings.Index(input, "m"); mi != -1 {
			m.line = line
			m.position = mi
			m.found = true
		}
		if pi := strings.Index(input, "p"); pi != -1 {
			p.line = line
			p.position = pi
			p.found = true
		}

		if m.found && p.found {
			break
		}
		line++
	}

	if !m.found || !p.found {
		return
	}

	ver := getVerticalSteps(int32(p.line-m.line), "UP", "DOWN")
	hor := getVerticalSteps(int32(p.position-m.position), "LEFT", "RIGHT")
	_, err := fmt.Fprint(os.Stdout, ver)
	if err != nil {
		return
	}
	_, err = fmt.Fprint(os.Stdout, hor)
	if err != nil {
		return
	}
}

func getVerticalSteps(n int32, txtf, txts string) string {
	var rst string
	if n == 0 {
		return rst
	}

	txt := txts
	if n < 0 {
		txt = txtf
	}
	n = int32(math.Abs(float64(n)))
	for i := int32(0); i < n; i++ {
		rst += fmt.Sprintf("%s\n", txt)
	}
	return rst
}
