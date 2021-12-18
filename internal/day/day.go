package day

import "io"

type Day[T, U any] interface {
	Run(r io.ReadSeeker) (T, U, error)
}

type Func[T, U any] func(r io.ReadSeeker) (T, U, error)

func (f Func[T, U]) Run(r io.ReadSeeker) (T, U, error) { return f(r) }
