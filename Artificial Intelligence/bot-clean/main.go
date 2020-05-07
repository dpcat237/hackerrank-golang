package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Bot struct {
	action   string
	line     int32
	position int32
}

type Dust struct {
	line     int32
	position int32
}

func main() {
	cleanBot()
}

func cleanBot() {
	in := bufio.NewReader(os.Stdin)
	var b Bot
	mDust := make(map[int32]map[int32]bool)

	line := int32(0)
	sline := int32(0)
	for {
		inputBt, _, err := in.ReadLine()
		input := string(inputBt)
		if err != nil {
			break
		}

		if sline == 0 {
			b.setPosition(input)
			sline++
			continue
		}

		lDust := findDust(input)
		if len(lDust) > 0 {
			mDust[line] = lDust
		}
		line++
	}

	if len(mDust) == 0 {
		return
	}

	if b.isCleanNow(mDust) {
		b.act()
		return
	}

	nextDust := b.getNextDust(mDust)
	if (Dust{}) == nextDust {
		return
	}

	b.moveToDust(nextDust)
	b.act()
}

func (b *Bot) moveToDust(d Dust) {
	if b.line != 0 && d.line < b.line {
		b.action = "UP"
		return
	}
	if d.line > b.line {
		b.action = "DOWN"
		return
	}
	if d.position < b.position {
		b.action = "LEFT"
		return
	}
	b.action = "RIGHT"
}

func (b Bot) act() {
	_, err := fmt.Fprint(os.Stdout, b.action)
	if err != nil {
		return
	}
}

func (b Bot) calculateEuclidean(dl, dp int32) int32 {
	return int32(math.Pow(float64(dl-b.line), 2) + math.Pow(float64(dp-b.position), 2))
}

func (b Bot) getNextDust(mDust map[int32]map[int32]bool) Dust {
	var dust Dust
	dist := int32(0)
	for l, lDust := range mDust {
		for p, _ := range lDust {
			dDist := b.calculateEuclidean(l, p)
			if dist == 0 {
				dist = dDist
				dust.set(l, p)
				continue
			}
			if dDist < dist {
				dist = dDist
				dust.set(l, p)
			}
		}
	}
	return dust
}

func (b *Bot) isCleanNow(mDust map[int32]map[int32]bool) bool {
	if lDust, ok := mDust[b.line]; ok {
		if _, ok := lDust[b.position]; ok {
			b.action = "CLEAN"
			return true
		}
	}
	return false
}

func (d *Dust) set(l, p int32) {
	d.line = l
	d.position = p
}

func (b *Bot) setPosition(line string) {
	words := strings.Fields(line)
	l, err := strconv.Atoi(words[0])
	if err != nil {
		os.Exit(1)
	}
	p, err := strconv.Atoi(words[1])
	if err != nil {
		os.Exit(1)
	}
	b.line = int32(l)
	b.position = int32(p)
}

func findDust(line string) map[int32]bool {
	ds := make(map[int32]bool)
	for i, p := range line {
		if p == 100 {
			ds[int32(i)] = true
		}
	}
	return ds
}
