package day

import "io"

type Day interface {
	Run(r io.ReadSeeker) (int, int, error)
}

type Func func(r io.ReadSeeker) (int, int, error)

func (f Func) Run(r io.ReadSeeker) (int, int, error) { return f(r) }
