package day8

import (
	"io"
	"strings"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

const (
	a = 1 << 6
	b = 1 << 5
	c = 1 << 4
	d = 1 << 3
	e = 1 << 2
	f = 1 << 1
	g = 1

	eight = a | b | c | d | e | f | g
)

type bitmask int

func newBitmask(str string) bitmask {
	i := 0
	for _, char := range str {
		i |= 1 << (char % 7)
	}
	return bitmask(i)
}

func (bm bitmask) Segments() int {
	return int((bm&a)>>6 + (bm&b)>>5 + (bm&c)>>4 + (bm&d)>>3 + (bm&e)>>2 + (bm&f)>>1 + (bm & g))
}

type data struct {
	input  []bitmask
	output []bitmask
}

func decoder(s string) (data, error) {
	parts := strings.Split(s, " | ")
	d := data{}
	for _, str := range strings.Split(parts[0], " ") {
		d.input = append(d.input, newBitmask(str))
	}
	for _, str := range strings.Split(parts[1], " ") {
		d.output = append(d.output, newBitmask(str))
	}
	return d, nil
}

func Run(r io.ReadSeeker) (int, int, error) {
	first, second := 0, 0
	scanner := reader.NewScanner[data](r, reader.NewDecoder(reader.NewBasicParser(decoder)))
	for scanner.Scan() {
		entry, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		bitmasks := identify(entry.input)
		sum := 0
		for _, v := range entry.output {
			if _, ok := isKnown(v.Segments()); ok {
				first++
			}
			vv := bitmasks[v]
			sum *= 10
			sum += vv
		}
		second += sum
	}
	return first, second, nil
}

func identify(input []bitmask) map[bitmask]int {
	known := map[int]bitmask{}
	unknown := make([]bitmask, 0, len(input))
	bitmasks := map[bitmask]int{}
	for _, v := range input {
		s := v.Segments()
		id, ok := isKnown(s)
		if !ok {
			unknown = append(unknown, v)
			continue
		}
		known[id] = v
		bitmasks[v] = id
	}
	for _, v := range unknown {
		seg := v.Segments()
		if r := known[7] | known[4]; seg == 6 && v&r == r {
			known[9] = v
			bitmasks[v] = 9
		} else if r := (v & known[7]) ^ known[7]; seg == 6 && r != 0 {
			known[6] = v
			bitmasks[v] = 6
		} else if seg == 5 && v&known[1] == known[1] {
			known[3] = v
			bitmasks[v] = 3
		} else if len(known) == 7 {
			break
		}
	}
	known[0] = ((known[4] | known[7]) ^ known[8]) | known[1] | (known[3] ^ known[4])
	known[5] = (known[4] & known[6]) | (known[3] ^ known[4])
	known[2] = (known[5] ^ known[0]) | (known[9] ^ known[4])
	bitmasks[known[0]] = 0
	bitmasks[known[5]] = 5
	bitmasks[known[2]] = 2
	return bitmasks
}

func isKnown(s int) (int, bool) {
	switch s {
	case 2:
		return 1, true
	case 4:
		return 4, true
	case 3:
		return 7, true
	case 7:
		return 8, true
	}
	return 0, false
}
