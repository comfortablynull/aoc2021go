package day4

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

type board struct {
	remaining int64
	rs, cs    [5]uint8
	values    map[int64][2]int
}

func newBoard() *board {
	return &board{values: map[int64][2]int{}}
}

func (b *board) hit(v int64) bool {
	xy, ok := b.values[v]
	if !ok {
		return false
	}
	delete(b.values, v)
	x, y := xy[0], xy[1]
	b.cs[x] += 1
	b.rs[y] += 1
	b.remaining -= v
	return b.rs[y] == 5 || b.cs[x] == 5
}

type boardParser struct {
	y int
	b *board
}

func newBoardParser() *boardParser {
	return &boardParser{b: newBoard()}
}

func (b *boardParser) Parse(s string) (continueReading bool, err error) {
	if s == "" {
		return false, nil
	}
	var row [5]int64
	sb := strings.Builder{}
	sv := s
	for k := range row {
		sb.Reset()
		sv = strings.TrimSpace(sv)
		for len(sv) != 0 && sv[0] != ' ' {
			sb.WriteByte(sv[0])
			sv = sv[1:]
		}
		v, err := strconv.ParseInt(sb.String(), 10, 64)
		if err != nil {
			return false, err
		}
		row[k] = v
	}
	for x, v := range row {
		b.b.values[v] = [2]int{x, b.y}
		b.b.remaining += v
	}
	b.y++
	return true, nil
}

func (b *boardParser) Result() *board {
	r := b.b
	b.b = newBoard()
	b.y = 0
	return r
}
func input(r io.ReadSeeker) []int64 {
	b := bufio.NewScanner(r)
	b.Scan()
	t := b.Text()
	split := strings.Split(t, ",")
	i := make([]int64, len(split))
	for k, v := range split {
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		i[k] = vv
	}
	return i
}

func Run(r io.ReadSeeker) (int, int, error) {
	inputs := input(r)
	if _, err := r.Seek(0, 0); err != nil {
		return 0, 0, err
	}
	scanner := reader.NewScanner(r, reader.NewDecoder[*board](newBoardParser()), reader.SkipLines(2))
	one, two := 0, 0
	win, lose := 0, 0
	for scanner.Scan() {
		b, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		v, n, ok := play(b, inputs)
		if !ok {
			continue
		}
		if n < win || win == 0 {
			one, win = v, n
		}
		if n > lose {
			two, lose = v, n
		}
	}
	return one, two, nil
}

func play(b *board, inputs []int64) (int, int, bool) {
	for k, v := range inputs {
		if b.hit(v) {
			return int(b.remaining * v), k, true
		}
	}
	return 0, len(inputs), false
}
