package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(recursiveSum())
}

func recursiveSum() int64 {
	var dgsStr string
	var rpt int64
	if _, err := fmt.Scanf("%s%d", &dgsStr, &rpt); err != nil {
		return 0
	}

	cut := 18
	if len(dgsStr) < cut {
		dgs, err := strconv.ParseInt(dgsStr, 10, 64)
		if err != nil {
			return 0
		}

		return sumDigits(sumDigits(dgs) * rpt)
	}

	var past bool
	for !past {
		dgsStr = sumDigitsHuge(dgsStr)
		if dgsStr == "" {
			return 0
		}

		if len(dgsStr) < cut {
			past = true
			continue
		}
	}

	dgsPart, err := strconv.ParseInt(dgsStr, 10, 64)
	if err != nil {
		return 0
	}

	return sumDigits(sumDigits(dgsPart) * rpt)
}

func sumDigitsHuge(dgsStr string) string {
	tt := int64(len(dgsStr))
	cut := int64(18)
	var oneSum string
	var cutFrom int64
	cutTo := cut
	var past bool
	for !past {
		dgsPart, err := strconv.ParseInt(dgsStr[cutFrom:cutTo], 10, 64)
		if err != nil {
			return ""
		}
		oneSum += fmt.Sprintf("%d", sumDigits(dgsPart))

		if tt == cutTo {
			past = true
		}

		if tt-cut >= cutTo {
			cutFrom += cut
			cutTo += cut
			continue
		}
		cutFrom = cutTo
		cutTo = tt
	}
	return oneSum
}

func sumDigits(n int64) int64 {
	if n < 10 {
		return n
	}

	var sum int64
	for n != 0 {
		sum += n % 10
		n = n / 10
	}
	return sumDigits(sum)
}
