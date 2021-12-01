package reader

import (
	"bufio"
	"errors"
	"io"
)

type Option interface {
	apply(s *bufio.Scanner)
}

type optFn func(s *bufio.Scanner)

func (o optFn) apply(s *bufio.Scanner) {
	o(s)
}

func SkipLines(n int) Option {
	return optFn(func(s *bufio.Scanner) {
		for i := 0; i < n; i++ {
			s.Scan()
		}
	})
}

var NoResultErr = errors.New("no result")

type Scanner[T any] struct {
	s *bufio.Scanner
	d Decoder[T]
}

func NewScanner[T any](r io.Reader, d Decoder[T], opts ...Option) *Scanner[T] {
	s := &Scanner[T]{s: bufio.NewScanner(r), d: d}
	for _, v := range opts {
		v.apply(s.s)
	}
	return s
}

func (i *Scanner[T]) Scan() bool {
	for i.s.Scan() && i.d.ReadLine(i.s.Text()) {
	}
	return i.d.HasState()
}

func (i *Scanner[T]) Result() (res T, err error) {
	if !i.d.HasState() {
		err = NoResultErr
		return
	}
	return i.d.Result(), i.d.Err()
}
