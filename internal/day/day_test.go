package day

import (
	"fmt"
	"os"
	"testing"

	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day10"
	"github.com/comfortablynull/aoc2021go/internal/day/day2"
	"github.com/comfortablynull/aoc2021go/internal/day/day3"
	"github.com/comfortablynull/aoc2021go/internal/day/day4"
	"github.com/comfortablynull/aoc2021go/internal/day/day5"
	"github.com/comfortablynull/aoc2021go/internal/day/day6"
	"github.com/comfortablynull/aoc2021go/internal/day/day7"
	"github.com/comfortablynull/aoc2021go/internal/day/day8"
	"github.com/comfortablynull/aoc2021go/internal/day/day9"
)

func TestDay(t *testing.T) {
	type test struct {
		label    string
		one, two int
		day      Day
	}
	tests := []test{
		{
			label: "one",
			one:   7,
			two:   5,
			day:   Func(day1.Run),
		},
		{
			label: "two",
			one:   150,
			two:   900,
			day:   Func(day2.Run),
		},
		{
			label: "three",
			one:   198,
			two:   230,
			day:   &day3.Runner{},
		},
		{
			label: "four",
			one:   4512,
			two:   1924,
			day:   Func(day4.Run),
		},
		{
			label: "five",
			one:   5,
			two:   12,
			day:   Func(day5.Run),
		},
		{
			label: "six",
			one:   5934,
			two:   26984457539,
			day:   Func(day6.Run),
		},
		{
			label: "seven",
			one:   37,
			two:   168,
			day:   Func(day7.Run),
		},
		{
			label: "eight",
			one:   26,
			two:   61229,
			day:   Func(day8.Run),
		},
		{
			label: "nine",
			one:   15,
			two:   1134,
			day:   Func(day9.Run),
		},
		{
			label: "ten",
			one:   26397,
			two:   288957,
			day:   Func(day10.Run),
		},
	}
	for _, tc := range tests {
		t.Run(tc.label, func(t *testing.T) {
			f, err := os.Open(fmt.Sprintf("testdata/%v.txt", tc.label))
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			one, two, err := tc.day.Run(f)
			if err != nil {
				t.Fatal(err)
			}
			if one != tc.one {
				t.Errorf("One wrong: expected: %v got: %v", tc.one, one)
			}
			if two != tc.two {
				t.Errorf("Two wrong: expected: %v got: %v", tc.two, two)
			}
		})
	}
}
