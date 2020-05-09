package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	text          string
	order         int32
	words         []string
	uniqueWordsID []string
	uniqueWords   []string
	answer        string
}

type Sentence struct {
	text      string
	wordsID   []string
	words     []string
	wordsHash string
}

type Sentences []Sentence

func main() {
	reply()
}

func reply() {
	in := bufio.NewReader(os.Stdin)
	var lines []string
	for {
		inputBt, _, err := in.ReadLine()
		input := string(inputBt)
		if err != nil {
			break
		}

		lines = append(lines, input)
	}

	if len(lines) == 0 {
		return
	}

	ses := splitSentences(lines[0])
	responses := strings.Split(lines[len(lines)-1], ";")
	questions := lines[1 : len(lines)-1]

	dic, dicRev, err := ses.createDictionary()
	if err != nil {
		return
	}
	qs := formalizeQuestions(questions, dic, dicRev)

	for _, q := range qs {
		q.findResponse(ses, responses)
		_, err := fmt.Fprint(os.Stdout, q.answer+"\n")
		if err != nil {
			return
		}
	}
}

func (q *Question) findResponse(ses Sentences, responses []string) {
	var max int32
	var seRes Sentence
	for _, se := range ses {
		rst := q.searchSentence(se.wordsHash)
		if max == 0 {
			max = rst
			seRes = se
		}
		if rst > max {
			max = rst
			seRes = se
		}
	}
	if max == 0 {
		return
	}

	var maxAns int32
	var ans string
	for _, resp := range responses {
		if strings.Contains(seRes.text, resp) {
			lenght := int32(len(resp))
			if maxAns == 0 {
				maxAns = lenght
				ans = resp
			}
			if lenght > maxAns {
				maxAns = lenght
				ans = resp
			}
		}
	}
	if maxAns == 0 {
		return
	}
	q.answer = ans
}

func (q Question) searchSentence(seHash string) int32 {
	words := int32(len(q.uniqueWordsID))
	return q.searchPart(seHash, 0, words, words)
}

func (q Question) searchPart(seHash string, from, to, words int32) int32 {
	if words == 0 {
		return 0
	}
	if strings.Contains(seHash, q.getPartHash(from, to)) {
		return words
	}

	from = from + 1
	to = to + 1
	if to > int32(len(q.uniqueWordsID)) {
		from = 0
		words = words - 1
		to = words
	}
	return q.searchPart(seHash, from, to, words)
}

func (q *Question) getPartHash(from, to int32) string {
	var hash string
	for _, wID := range q.uniqueWordsID[from:to] {
		hash += wID
	}
	return hash
}

func (se Sentence) countrySame(wsIDs []string) int32 {
	var c int32
	for _, wsID := range wsIDs {
		for _, wID := range se.wordsID {
			if wsID == wID {
				c++
				break
			}
		}
	}

	return c
}

func formalizeQuestions(questionsTxt []string, dic map[string]string, dicRev map[string]string) []Question {
	var qs []Question
	for key, queTxt := range questionsTxt {
		q := Question{
			text:  queTxt,
			order: int32(key),
		}
		q.setWords(dic)
		qs = append(qs, q)
	}

	for k := range qs {
		qs[k].setUnique(qs, dicRev)
	}
	return qs
}

func (q *Question) setWords(dic map[string]string) {
	for _, word := range getWords(q.text) {
		word := strings.ToLower(word)
		if wID, ok := dic[word]; ok {
			q.words = append(q.words, wID)
		}
	}
}

func (que Question) isWordUnique(wID string, qs []Question) bool {
	for _, q := range qs {
		if q.order == que.order {
			continue
		}
		for _, qwID := range q.words {
			if qwID == wID {
				return false
			}
		}
	}
	return true
}

func (q *Question) setUnique(qs []Question, dicRev map[string]string) {
	for _, wID := range q.words {
		if q.isWordUnique(wID, qs) {
			q.uniqueWordsID = append(q.uniqueWordsID, wID)
			q.uniqueWords = append(q.uniqueWords, dicRev[wID])
		}
	}
}

func splitSentences(text string) Sentences {
	var ses []Sentence
	if strings.Contains(text, "!") {
		text = strings.ReplaceAll(text, "!", ".")
	}
	for _, sTxt := range strings.Split(text, ".") {
		se := Sentence{
			text: sTxt,
		}
		ses = append(ses, se)
	}
	return ses
}

func (se *Sentence) setHash() {
	for _, wId := range se.wordsID {
		se.wordsHash += wId
	}
}

func (ses *Sentences) createDictionary() (map[string]string, map[string]string, error) {
	dic := make(map[string]string)
	dicRev := make(map[string]string)
	c := int32(1)
	for sek, se := range *ses {
		for _, word := range getWords(se.text) {
			word = strings.TrimSpace(strings.ToLower(word))
			dID, ok := dic[word]
			if ok {
				(*ses)[sek].wordsID = append((*ses)[sek].wordsID, dID)
				(*ses)[sek].words = append((*ses)[sek].words, word)
				continue
			}
			cID := toID(c)
			dic[word] = cID
			dicRev[cID] = word
			(*ses)[sek].wordsID = append((*ses)[sek].wordsID, cID)
			(*ses)[sek].words = append((*ses)[sek].words, word)
			c++
		}
		(*ses)[sek].setHash()
	}
	return dic, dicRev, nil
}

func toID(num int32) string {
	return fmt.Sprintf("000000%d", num)
}

func getWords(sentence string) []string {
	return strings.Split(strip(sentence), " ")
}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}
