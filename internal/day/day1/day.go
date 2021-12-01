package day1

import (
	"io"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

func Run(r io.ReadSeeker) (int, int, error) {
	scanner := reader.NewScanner[int](r, reader.NewIntDecoder(10, 64))
	scanner.Scan()
	last, err := scanner.Result()
	if err != nil {
		return 0, 0, err
	}
	buff := make([]int, 3)
	buff[0] = last
	sum := last
	k := 1
	first, second := 0, 0
	for scanner.Scan() {
		v, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		if v > last {
			first++
		}
		last = v
		idx := k % 3
		sum, previous := sum-buff[idx]+v, sum
		buff[idx] = v
		// TODO: if I have some time later do something better than k >= 3 for every check
		if k >= 3 && sum > previous {
			second++
		}
		k++
	}
	return first, second, nil
}
