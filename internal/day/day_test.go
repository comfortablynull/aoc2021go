package day

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day10"
	"github.com/comfortablynull/aoc2021go/internal/day/day11"
	"github.com/comfortablynull/aoc2021go/internal/day/day12"
	"github.com/comfortablynull/aoc2021go/internal/day/day2"
	"github.com/comfortablynull/aoc2021go/internal/day/day3"
	"github.com/comfortablynull/aoc2021go/internal/day/day4"
	"github.com/comfortablynull/aoc2021go/internal/day/day5"
	"github.com/comfortablynull/aoc2021go/internal/day/day6"
	"github.com/comfortablynull/aoc2021go/internal/day/day7"
	"github.com/comfortablynull/aoc2021go/internal/day/day8"
	"github.com/comfortablynull/aoc2021go/internal/day/day9"
	"github.com/google/go-cmp/cmp"
)

type TestCase interface {
	Run(t *testing.T, f io.ReadSeeker)
}

type testcase[T, U comparable] struct {
	one T
	two U
	day Day[T, U]
}

func (tc testcase[T, U]) Run(t *testing.T, f io.ReadSeeker) {
	one, two, err := tc.day.Run(f)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(tc.one, one); diff != "" {
		t.Errorf("One Wrong:\n%v", diff)
	}
	if diff := cmp.Diff(tc.two, two); diff != "" {
		t.Errorf("Two wrong:\n%v", diff)
	}
}

func TestDay(t *testing.T) {
	type test struct {
		label string
		TestCase
	}
	tests := []test{
		{
			label: "1",
			TestCase: testcase[int, int]{
				one: 7,
				two: 5,
				day: Func[int, int](day1.Run)},
		},
		{
			label: "2",
			TestCase: testcase[int, int]{
				one: 150,
				two: 900,
				day: Func[int, int](day2.Run)},
		},
		{
			label: "3",
			TestCase: testcase[int, int]{
				one: 198,
				two: 230,
				day: &day3.Runner{}},
		},
		{
			label: "4",
			TestCase: testcase[int, int]{
				one: 4512,
				two: 1924,
				day: Func[int, int](day4.Run)},
		},
		{
			label: "5",
			TestCase: testcase[int, int]{
				one: 5,
				two: 12,
				day: Func[int, int](day5.Run)},
		},
		{
			label: "6",
			TestCase: testcase[int, int]{
				one: 5934,
				two: 26984457539,
				day: Func[int, int](day6.Run)},
		},
		{
			label: "7",
			TestCase: testcase[int, int]{
				one: 37,
				two: 168,
				day: Func[int, int](day7.Run)},
		},
		{
			label: "8",
			TestCase: testcase[int, int]{
				one: 26,
				two: 61229,
				day: Func[int, int](day8.Run)},
		},
		{
			label: "9",
			TestCase: testcase[int, int]{
				one: 15,
				two: 1134,
				day: Func[int, int](day9.Run)},
		},
		{
			label: "10",
			TestCase: testcase[int, int]{
				one: 26397,
				two: 288957,
				day: Func[int, int](day10.Run)},
		},
		{
			label: "11",
			TestCase: testcase[int, int]{
				one: 1656,
				two: 195,
				day: Func[int, int](day11.Run)},
		},
		{
			label: "12",
			TestCase: testcase[int, int]{
				one: 226,
				two: 3509,
				day: Func[int, int](day12.Run)},
		},
	}
	for _, tc := range tests {
		t.Run(tc.label, func(t *testing.T) {
			f, err := os.Open(fmt.Sprintf("testdata/%v.txt", tc.label))
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			tc.Run(t, f)
		})
	}
}
