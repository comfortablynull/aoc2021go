package day

import (
	"io"

	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day10"
	"github.com/comfortablynull/aoc2021go/internal/day/day11"
	"github.com/comfortablynull/aoc2021go/internal/day/day12"
	"github.com/comfortablynull/aoc2021go/internal/day/day13"
	"github.com/comfortablynull/aoc2021go/internal/day/day2"
	"github.com/comfortablynull/aoc2021go/internal/day/day3"
	"github.com/comfortablynull/aoc2021go/internal/day/day4"
	"github.com/comfortablynull/aoc2021go/internal/day/day5"
	"github.com/comfortablynull/aoc2021go/internal/day/day6"
	"github.com/comfortablynull/aoc2021go/internal/day/day7"
	"github.com/comfortablynull/aoc2021go/internal/day/day8"
	"github.com/comfortablynull/aoc2021go/internal/day/day9"
)

type Day[T, U any] interface {
	Run(r io.ReadSeeker) (T, U, error)
}

type Func[T, U any] func(r io.ReadSeeker) (T, U, error)

func (f Func[T, U]) Run(r io.ReadSeeker) (T, U, error) { return f(r) }

func One() Day[int, int] {
	return Func[int, int](day1.Run)
}
func Two() Day[int, int] {
	return Func[int, int](day2.Run)
}
func Three() Day[int, int] {
	return &day3.Runner{}
}
func Four() Day[int, int] {
	return Func[int, int](day4.Run)
}
func Five() Day[int, int] {
	return Func[int, int](day5.Run)
}
func Six() Day[int, int] {
	return Func[int, int](day6.Run)
}
func Seven() Day[int, int] {
	return Func[int, int](day7.Run)
}
func Eight() Day[int, int] {
	return Func[int, int](day8.Run)
}
func Nine() Day[int, int] {
	return Func[int, int](day9.Run)
}
func Ten() Day[int, int] {
	return Func[int, int](day10.Run)
}
func Eleven() Day[int, int] {
	return Func[int, int](day11.Run)
}
func Twelve() Day[int, int] {
	return Func[int, int](day12.Run)
}
func Thirteen() Day[int, string] {
	return Func[int, string](day13.Run)
}
