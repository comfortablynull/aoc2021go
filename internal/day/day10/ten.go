package day10

import (
	"bufio"
	"io"
	"sort"
)

func Run(r io.ReadSeeker) (int, int, error) {
	scanner := bufio.NewScanner(r)
	first := 0
	var scores []int
START:
	for scanner.Scan() {
		var stack []uint8
		line := scanner.Text()
		for k := range line {
			r := line[k]
			if r == '{' || r == '[' || r == '<' || r == '(' {
				stack = append(stack, r)
				continue
			}
			if l := stack[len(stack)-1]; r-1 == l || r-2 == l {
				stack = stack[:len(stack)-1]
				continue
			}
			switch r {
			case ')':
				first += 3
			case ']':
				first += 57
			case '}':
				first += 1197
			case '>':
				first += 25137
			}
			goto START
		}
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			r := stack[i]
			score *= 5
			switch r {
			case '(':
				score += 1
			case '[':
				score += 2
			case '{':
				score += 3
			case '<':
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return first, scores[len(scores)/2], nil
}
