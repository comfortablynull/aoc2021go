package day6

import (
	"bufio"
	"io"
	"strings"
)

const (
	firstIterations = 80
	maxIterations   = 256
	resetStage      = 6
	nStages         = 8
)

func input(r io.ReadSeeker) []int {
	b := bufio.NewScanner(r)
	b.Scan()
	line := b.Text()
	entries := strings.Split(line, ",")
	ret := make([]int, len(entries))
	for k, v := range entries {
		ret[k] = int(v[0] - '0')
	}
	return ret
}

func Run(r io.ReadSeeker) (int, int, error) {
	state := input(r)
	first, second := len(state), len(state)
	stages := make([]int, nStages+1)
	for _, v := range state {
		stages[v]++
	}
	for i := 0; i < maxIterations; i++ {
		n := stages[0]
		if i < firstIterations {
			first += n
		}
		second += n
		for j := 1; j < len(stages); j++ {
			stages[j-1] += stages[j]
			stages[j] = 0
		}
		stages[0] -= n
		stages[resetStage] += n
		stages[nStages] += n
	}
	return first, second, nil
}
