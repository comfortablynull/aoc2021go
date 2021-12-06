package day

import (
	"fmt"
	"os"
	"testing"

	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day2"
	"github.com/comfortablynull/aoc2021go/internal/day/day3"
	"github.com/comfortablynull/aoc2021go/internal/day/day4"
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
