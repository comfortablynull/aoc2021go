package day11

import (
	"io"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

const (
	maxIterations = 100
	flash         = 10
)

type point struct {
	v, i int
}

func pointDec(s string) ([10]point, error) {
	var r [10]point
	for k := range s {
		r[k].v = int(s[k] - '0')
	}
	return r, nil
}

func Run(r io.ReadSeeker) (int, int, error) {
	var first int
	var grid [10][10]point
	scanner := reader.NewScanner(r, reader.NewDecoder(reader.NewBasicParser(pointDec)))
	for k := 0; scanner.Scan() && k < 10; k++ {
		r, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		grid[k] = r
	}
	for i := 1; ; i++ {
		count := 0
		for y, row := range grid {
			for x := range row {
				count += inc(i, x, y, &grid)
			}
		}
		if i <= maxIterations {
			first += count
		}
		if count == 100 {
			return first, i, nil
		}
	}
}

func inc(i, x, y int, grid *[10][10]point) int {
	if y < 0 || y > 9 || x < 0 || x > 9 || grid[y][x].i == i {
		return 0
	}
	grid[y][x].v++
	if grid[y][x].v%flash != 0 {
		return 0
	}
	grid[y][x].i = i
	return 1 + inc(i, x-1, y+1, grid) +
		inc(i, x, y+1, grid) +
		inc(i, x+1, y+1, grid) +
		inc(i, x, y-1, grid) +
		inc(i, x-1, y-1, grid) +
		inc(i, x+1, y-1, grid) +
		inc(i, x+1, y, grid) +
		inc(i, x-1, y, grid)
}
