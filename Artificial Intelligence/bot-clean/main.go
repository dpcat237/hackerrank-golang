package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bot struct {
	action       string
	line         int32
	position     int32
	positionated bool
	priority     string
	total        int32
}

func main() {
	cleanBot()
}

func cleanBot() {
	in := bufio.NewReader(os.Stdin)
	var b Bot
	mDust := make(map[int32]map[int32]bool)

	line := int32(0)
	for {
		inputBt, _, err := in.ReadLine()
		input := string(inputBt)
		if err != nil {
			break
		}

		if line == 0 && !b.positionated {
			b.positionate(input)
			continue
		}

		b.setTotal(input)
		mDust[line] = findDust(input)

		if b.line == line-1 && len(mDust[line-1]) > 0 {
			b.cleanLine(mDust)
			break
		}

		line++
	}

	b.cleanLine(mDust)
	if b.action == "" {
		b.changeLine(mDust)
	}
	if b.action == "" {
		return
	}

	_, err := fmt.Fprint(os.Stdout, b.action)
	if err != nil {
		return
	}
}

func (b *Bot) changeLine(mDust map[int32]map[int32]bool) {
	if b.line == 0 {
		if b.isDustThere(1, mDust) {
			b.action = "DOWN"
		}
		return
	}

	if b.line == (b.total - 1) {
		if b.isDustThere(b.total-2, mDust) {
			b.action = "UP"
		}
		return
	}

	if b.line < (b.total / 2) {
		if b.isDustThere(b.line-1, mDust) {
			b.action = "UP"
			return
		}
		if b.isDustThere(b.line+1, mDust) {
			b.action = "DOWN"
		}
		return
	}

	if b.isDustThere(b.line+1, mDust) {
		b.action = "DOWN"
		return
	}
	if b.isDustThere(b.line-1, mDust) {
		b.action = "UP"
	}
}

func (b *Bot) checkDustInNextLine(mDust map[int32]map[int32]bool) {
	if b.priority == "DOWN" && b.line < b.total-1 {
		if _, ok := mDust[b.line+1][b.position]; ok {
			b.action = "DOWN"
		}
	}
	if b.line > 0 {
		if _, ok := mDust[b.line-1][b.position]; ok {
			b.action = "UP"
		}
	}
}

func (b *Bot) cleanLine(mDust map[int32]map[int32]bool) {
	ds := mDust[b.line]
	if len(ds) == 0 {
		return
	}
	if _, ok := ds[b.position]; ok {
		b.action = "CLEAN"
		return
	}

	b.checkDustInNextLine(mDust)
	if b.action != "" {
		return
	}

	left := b.getSteps(b.position-1, 1, ds)
	right := b.getSteps(b.position+1, 1, ds)

	if left == 0 {
		b.action = "RIGHT"
		return
	}

	if right == 0 {
		b.action = "LEFT"
		return
	}

	if left == right {
		b.action = "LEFT"
		return
	}

	if left < right {
		b.action = "LEFT"
		return
	}
	b.action = "RIGHT"
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

func (b *Bot) getSteps(i, done int32, ds map[int32]bool) int32 {
	if _, ok := ds[i]; ok {
		return done
	}

	if i <= 0 || i >= (b.total-1) {
		return 0
	}

	done++
	if i < b.position {
		return b.getSteps(i-1, done, ds)
	}
	return b.getSteps(i+1, done, ds)
}

func (b Bot) isDustThere(l int32, mDust map[int32]map[int32]bool) bool {
	if len(mDust[l]) > 0 {
		return true
	}

	if l == 0 || l == b.total-1 {
		return false
	}
	if l < b.line {
		return b.isDustThere(l-1, mDust)
	}
	return b.isDustThere(l+1, mDust)
}

func (b *Bot) positionate(line string) {
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
	b.positionated = true
}

func (b *Bot) setPriority() {
	if b.line < (b.total / 2) {
		b.priority = "DOWN"
		return
	}
	b.priority = "UP"
}

func (b *Bot) setTotal(line string) {
	if b.total == 0 {
		b.total = int32(len(line))
		b.setPriority()
	}
}
