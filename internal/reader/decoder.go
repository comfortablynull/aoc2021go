package reader

import "strconv"

type ParserFunc[T any] func(input string) (result T, err error)

type Parser[T any] interface {
	Parse(s string) (continueReading bool, err error)
	Result() T
}

type basicParser[T any] struct {
	val T
	fn  ParserFunc[T]
}

func NewBasicParser[T any](fn ParserFunc[T]) Parser[T] {
	return &basicParser[T]{fn: fn}
}

func (b *basicParser[T]) Parse(s string) (continueReading bool, err error) {
	v, err := b.fn(s)
	b.val = v
	return false, err
}

func (b *basicParser[T]) Result() T {
	return b.val
}

type Decoder[T any] interface {
	ReadLine(string) bool
	HasState() bool
	Result() T
	Err() error
}

type decoder[T any] struct {
	state  bool
	err    error
	parser Parser[T]
}

func (b *decoder[T]) ReadLine(s string) bool {
	cont, err := b.parser.Parse(s)
	b.state = true
	if err != nil {
		b.err = err
		return false
	}
	return cont
}

func (b decoder[T]) HasState() bool {
	return b.state
}

func (b *decoder[T]) Result() T {
	b.state = false
	return b.parser.Result()
}

func (b *decoder[T]) Err() error {
	b.state = false
	return b.err
}

func NewDecoder[T any](parser Parser[T]) Decoder[T] {
	return &decoder[T]{parser: parser}
}

func NewIntDecoder(base, bits int) Decoder[int] {
	return NewDecoder[int](NewBasicParser[int](func(s string) (int, error) {
		v, err := strconv.ParseInt(s, base, bits)
		return int(v), err
	}))
}

func NewInt64Decoder(base, bits int) Decoder[int64] {
	return NewDecoder[int64](NewBasicParser[int64](func(s string) (int64, error) {
		v, err := strconv.ParseInt(s, base, bits)
		return v, err
	}))
}

func NewStringDecoder() Decoder[string] {
	return NewDecoder[string](NewBasicParser[string](func(s string) (string, error) {
		return s, nil
	}))
}

func NewQuickIntSliceDecoder() Decoder[[]int] {
	return NewDecoder[[]int](NewBasicParser[[]int](func(s string) ([]int, error) {
		m := make([]int, len(s))
		for k := range s {
			m[k] = int(s[k] - '0')
		}
		return m, nil
	}))
}
