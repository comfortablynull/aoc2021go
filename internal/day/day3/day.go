package day3

import (
	"io"
	"strconv"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

const (
	n0 = '0'
	n1 = '1'
)

type Runner struct {
	max     uint8
	initial []string
}

func (r2 *Runner) Run(r io.ReadSeeker) (int, int, error) {
	one, err := r2.one(r)
	if err != nil {
		return 0, 0, err
	}
	two, err := r2.two()
	return one, two, err
}

func (r2 *Runner) one(r io.ReadSeeker) (int, error) {
	scanner := reader.NewScanner[string](r, reader.NewStringDecoder())
	var counts []int
	count := 0
	for scanner.Scan() {
		v, err := scanner.Result()
		if err != nil {
			return 0, err
		}
		count++
		if lc, lv := len(counts), len(v); lv > lc {
			counts = append(counts, make([]int, lv-lc)...)
		}
		for i := 0; i < len(v); i++ {
			counts[i] += int(v[i] - n0)
		}
		r2.initial = append(r2.initial, v)
	}
	if r := count - counts[0]; counts[0] > r {
		r2.max = n1
	} else {
		r2.max = n0
	}
	gamma, epsilon := 0, 0
	for _, v := range counts {
		// shift one to the left and store
		gamma <<= 1
		epsilon <<= 1
		// set last bit to 1
		if r := count - v; v > r {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	return gamma * epsilon, nil
}

func (r2 *Runner) two() (int, error) {
	oxy, err := filter(r2.initial, func(c int, l int) uint8 {
		if rr := l - c; c >= rr {
			return n1
		}
		return n0
	})
	if err != nil {
		return 0, err
	}
	co2, err := filter(r2.initial, func(c int, l int) uint8 {
		if rr := l - c; c >= rr {
			return n0
		}
		return n1
	})
	if err != nil {
		return 0, err
	}
	return co2 * oxy, nil
}

func filter(initial []string, fn func(int, int) uint8) (int, error) {
	state := append(make([]string, 0, len(initial)), initial...)
	for i, l := 0, len(state); l != 1; i, l = i+1, len(state) {
		c := 0
		for _, v := range state {
			c += int(v[i] - n0)
		}
		target := fn(c, l)
		n := 0
		for _, v := range state {
			if v[i] != target {
				continue
			}
			state[n] = v
			n++
		}
		state = state[:n]
	}
	v, err := strconv.ParseInt(state[0], 2, 64)
	return int(v), err
}
