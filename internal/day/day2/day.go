package day2

import (
	"io"
	"strconv"
	"strings"

	"github.com/comfortablynull/aoc2021go/internal/reader"
)

type Point struct {
	X, Y int
}

func PointDecoder(s string) (Point, error) {
	p := Point{}
	parts := strings.Split(s, " ")
	val, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return p, err
	}
	n := int(val)
	switch parts[0] {
	case "forward":
		p.X = n
	case "down":
		p.Y = n
	case "up":
		p.Y = -1 * n
	}
	return p, nil
}

func Run(r io.ReadSeeker) (int, int, error) {
	scanner := reader.NewScanner[Point](r, reader.NewDecoder(reader.NewBasicParser(PointDecoder)))
	var p1, p2 Point
	var aim int
	for scanner.Scan() {
		p, err := scanner.Result()
		if err != nil {
			return 0, 0, err
		}
		p1.X += p.X
		p1.Y += p.Y
		p2.X += p.X
		aim += p.Y
		p2.Y += aim * p.X
	}
	return p1.X * p1.Y, p2.X * p2.Y, nil
}
