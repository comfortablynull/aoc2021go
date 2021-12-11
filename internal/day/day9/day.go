package day9

import (
	"io"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

func Run(r io.ReadSeeker) (int, int, error) {
	first := 0
	scanner := reader.NewScanner[[]int](r, reader.NewQuickIntSliceDecoder())
	var grid [][]int
	for scanner.Scan() {
		row, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		grid = append(grid, row)
	}
	ym := len(grid) - 1
	xm := len(grid[0]) - 1
	var largest [3]int
	for y, row := range grid {
		for x, col := range row {
			if y != 0 && grid[y-1][x] <= col {
				continue
			} else if x != 0 && grid[y][x-1] <= col {
				continue
			} else if y != ym && grid[y+1][x] <= col {
				continue
			} else if x != xm && grid[y][x+1] <= col {
				continue
			}
			first += col + 1
			vv := sumBasin(x, y, grid)
			if vv > largest[2] {
				largest[0], largest[1] = largest[1], largest[2]
				largest[2] = vv
			} else if vv > largest[1] {
				largest[0] = largest[1]
				largest[1] = vv
			} else if vv > largest[0] {
				largest[0] = vv
			}
		}
	}
	return first, largest[0] * largest[1] * largest[2], nil
}

func sumBasin(x, y int, grid [][]int) int {
	if x < 0 || y < 0 || y == len(grid) || x == len(grid[y]) || grid[y][x] == 9 {
		return 0
	}
	grid[y][x] = 9
	return 1 + sumBasin(x+1, y, grid) + sumBasin(x-1, y, grid) + sumBasin(x, y-1, grid) + sumBasin(x, y+1, grid)
}
