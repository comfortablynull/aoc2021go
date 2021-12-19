package day

import (
	"fmt"
	"io"
	"os"
	"testing"

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
				day: One(),
			},
		},
		{
			label: "2",
			TestCase: testcase[int, int]{
				one: 150,
				two: 900,
				day: Two(),
			},
		},
		{
			label: "3",
			TestCase: testcase[int, int]{
				one: 198,
				two: 230,
				day: Three(),
			},
		},
		{
			label: "4",
			TestCase: testcase[int, int]{
				one: 4512,
				two: 1924,
				day: Four(),
			},
		},
		{
			label: "5",
			TestCase: testcase[int, int]{
				one: 5,
				two: 12,
				day: Five(),
			},
		},
		{
			label: "6",
			TestCase: testcase[int, int]{
				one: 5934,
				two: 26984457539,
				day: Six(),
			},
		},
		{
			label: "7",
			TestCase: testcase[int, int]{
				one: 37,
				two: 168,
				day: Seven(),
			},
		},
		{
			label: "8",
			TestCase: testcase[int, int]{
				one: 26,
				two: 61229,
				day: Eight(),
			},
		},
		{
			label: "9",
			TestCase: testcase[int, int]{
				one: 15,
				two: 1134,
				day: Nine(),
			},
		},
		{
			label: "10",
			TestCase: testcase[int, int]{
				one: 26397,
				two: 288957,
				day: Ten(),
			},
		},
		{
			label: "11",
			TestCase: testcase[int, int]{
				one: 1656,
				two: 195,
				day: Eleven(),
			},
		},
		{
			label: "12",
			TestCase: testcase[int, int]{
				one: 226,
				two: 3509,
				day: Twelve(),
			},
		},
		{
			label: "13",
			TestCase: testcase[int, string]{
				one: 17,
				two: "#####\n#...#\n#...#\n#...#\n#####\n.....\n.....",
				day: Thirteen(),
			},
		}, {
			label: "14",
			TestCase: testcase[int, int]{
				one: 1588,
				two: 2188189693529,
				day: Fourteen(),
			},
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
