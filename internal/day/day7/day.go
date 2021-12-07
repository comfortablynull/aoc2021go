package day7

import (
	"bufio"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

func input(r io.ReadSeeker) ([]int, error) {
	b := bufio.NewScanner(r)
	b.Scan()
	line := b.Text()
	entries := strings.Split(line, ",")
	ret := make([]int, len(entries))
	for k, v := range entries {
		vv, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ret[k] = int(vv)
	}
	return ret, nil
}
func Run(r io.ReadSeeker) (int, int, error) {
	inputs, err := input(r)
	if err != nil {
		return 0, 0, err
	}
	sort.Ints(inputs)
	mid := inputs[len(inputs)/2]
	avg := float64(0)
	for _, v := range inputs {
		avg += float64(v)
	}
	avg /= float64(len(inputs))
	fl, cl := int(math.Floor(avg)), int(math.Ceil(avg))
	first, second := 0, 0
	s1, s2 := 0, 0
	for _, v := range inputs {
		first += int(math.Abs(float64(v - mid)))
		d1, d2 := int(math.Abs(float64(v-fl))), int(math.Abs(float64(v-cl)))
		s1 += d1 * (d1 + 1) / 2
		s2 += d2 * (d2 + 1) / 2
	}
	second = s1
	if s2 < second {
		second = s2
	}
	return first, second, nil
}
