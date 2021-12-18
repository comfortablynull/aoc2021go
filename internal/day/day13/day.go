package day13

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

type row struct {
	count  int
	column []bool
}

type Board struct {
	paper []row
	folds [][2]int
}

func (b *Board) Print() string {
	sb := strings.Builder{}
	for k, v := range b.paper {
		if k > 0 {
			sb.WriteRune('\n')
		}
		for _, vv := range v.column {
			if vv {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
	}
	return sb.String()
}

func (b *Board) Fold() (int, int) {
	var c1, c2 int
	for k, v := range b.folds {
		var c int
		if v[0] == 0 {
			c = b.y(v[1])
		} else {
			c = b.x(v[0])
		}
		if k == 0 {
			c1 = c
		}
		c2 = c
	}
	return c1, c2
}

func (b *Board) x(x int) int {
	var c int
	for ky, r := range b.paper {
		for k, v := range r.column[x+1:] {
			kx := x - k - 1
			if !v {
				continue
			}
			if b.paper[ky].column[kx] {
				b.paper[ky].count--
			}
			b.paper[ky].column[kx] = true
		}
		c += b.paper[ky].count
		b.paper[ky].column = b.paper[ky].column[:x]
	}
	return c
}

func (b *Board) y(y int) int {
	var c int
	for k, r := range b.paper[y+1:] {
		ky := y - k - 1
		for kx, v := range r.column {
			if !v || b.paper[ky].column[kx] {
				continue
			}
			b.paper[ky].column[kx] = true
			b.paper[ky].count++
		}
		c += b.paper[ky].count
	}
	b.paper = b.paper[:y]
	return c
}

type boardParser struct {
	board *Board
	xm    int
	f     func(s string) error
}

func newBoardParser() reader.Parser[*Board] {
	return &boardParser{board: &Board{}}
}

func (b *boardParser) Parse(s string) (continueReading bool, err error) {
	if s == "" {
		b.f = b.fold
		m := b.xm + 1
		for k, v := range b.board.paper {
			if l := m - len(v.column); l > 0 {
				b.board.paper[k].column = append(v.column, make([]bool, l)...)
			}
		}
		return true, nil
	}
	if b.f == nil {
		b.f = b.row
	}
	return true, b.f(s)
}

func (b *boardParser) Result() *Board {
	return b.board
}

func (b *boardParser) row(s string) error {
	xy := strings.Split(s, ",")
	x64, err := strconv.ParseInt(xy[0], 10, 64)
	if err != nil {
		return err
	}
	y64, err := strconv.ParseInt(xy[1], 10, 64)
	if err != nil {
		return err
	}
	x, y := int(x64), int(y64)
	if l := y - len(b.board.paper) + 1; l > 0 {
		b.board.paper = append(b.board.paper, make([]row, l)...)
	}
	if l := x - len(b.board.paper[y].column) + 1; l > 0 {
		b.board.paper[y].column = append(b.board.paper[y].column, make([]bool, l)...)
	}
	if x > b.xm {
		b.xm = x
	}
	b.board.paper[y].count += 1
	b.board.paper[y].column[x] = true
	return nil
}

func (b boardParser) fold(s string) error {
	cv := strings.Split(s[11:], "=")
	v, err := strconv.ParseInt(cv[1], 10, 64)
	if err != nil {
		return err
	}
	switch cv[0] {
	case "x":
		b.board.folds = append(b.board.folds, [2]int{int(v), 0})
	case "y":
		b.board.folds = append(b.board.folds, [2]int{0, int(v)})
	default:
		return errors.New("invalid fold")
	}
	return nil
}

func Run(r io.ReadSeeker) (int, string, error) {
	var first int
	var second string
	var err error
	var b *Board
	scanner := reader.NewScanner[*Board](r, reader.NewDecoder(newBoardParser()))
	if !scanner.Scan() {
		return first, second, errors.New("oh no")
	}
	b, err = scanner.Result()
	if err != nil {
		return first, second, err
	}
	first, _ = b.Fold()
	return first, b.Print(), err
}
