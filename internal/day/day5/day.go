package day5

import (
	"io"
	"strconv"
	"strings"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

type point struct {
	x, y int
}

func Range(input string) ([2]point, error) {
	parts := strings.Split(input, " -> ")
	first, err := Point(parts[0])
	if err != nil {
		return [2]point{}, err
	}
	second, err := Point(parts[1])
	return [2]point{first, second}, nil
}

func Point(s string) (point, error) {
	parts := strings.Split(s, ",")
	x, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return point{}, err
	}
	y, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return point{}, err
	}
	return point{x: int(x), y: int(y)}, nil
}

func Run(r io.ReadSeeker) (int, int, error) {
	var b1, b2 [][]int
	scanner := reader.NewScanner(r, reader.NewDecoder(reader.NewBasicParser(Range)))
	var first, second int
	for scanner.Scan() {
		entry, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		dx, mx, maxX := mod(entry[0].x, entry[1].x)
		dy, my, maxY := mod(entry[0].y, entry[1].y)
		if r := 1 + maxY - len(b1); r > 0 {
			p := make([][]int, r)
			b1 = append(b1, p...)
			b2 = append(b2, p...)
		}
		for x, y := entry[0].x, entry[0].y; x != mx || y != my; x, y = x+dx, y+dy {
			if r := maxX + 1 - len(b1[y]); r > 0 {
				p := make([]int, r)
				b1[y] = append(b1[y], p...)
				b2[y] = append(b2[y], p...)
			}
			if dx == 0 || dy == 0 {
				b1[y][x]++
				if b1[y][x] == 2 {
					first++
				}
			}
			b2[y][x]++
			if b2[y][x] == 2 {
				second++
			}
		}
	}
	return first, second, nil
}

func mod(a, b int) (int, int, int) {
	if a == b {
		return 0, b, b
	}
	if a > b {
		return -1, b - 1, a
	}
	return 1, b + 1, b
}
