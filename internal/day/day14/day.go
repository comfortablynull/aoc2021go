package day14

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func key(a, b uint8) [2]uint8 {
	return [2]uint8{a, b}
}

func Run(r io.ReadSeeker) (int, int, error) {
	var first, second int
	s := bufio.NewScanner(r)
	patterns := map[[2]uint8]uint8{}
	if !s.Scan() {
		return first, second, errors.New("can't read")
	}
	template := s.Text()
	s.Scan()
	for s.Scan() {
		spl := strings.Split(s.Text(), " -> ")
		patterns[key(spl[0][0], spl[0][1])] = spl[1][0]
	}
	pairs := map[[2]uint8]int{}
	counts := map[uint8]int{template[0]: 1}
	for i := 1; i < len(template); i++ {
		pairs[key(template[i-1], template[i])]++
		counts[template[i]]++
	}
	for i := 0; i < 40; i++ {
		np := map[[2]uint8]int{}
		for k, v := range pairs {
			c, ok := patterns[k]
			if !ok {
				continue
			}
			np[key(k[0], c)] += v
			np[key(c, k[1])] += v
			counts[c] += v
		}
		pairs = np
		if i == 9 {
			first = minMax(counts)
		}
	}
	second = minMax(counts)
	return first, second, nil
}

func minMax(counts map[uint8]int) int {
	min, max := -1, 0
	for _, v := range counts {
		if v < min || min == -1 {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}
