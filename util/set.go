package util

type Set[T comparable] map[T]bool

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s *Set[T]) Add(value T) {
	(*s)[value] = true
}

func (s *Set[T]) Remove(value T) {
	delete(*s, value)
}

func (s *Set[T]) Contains(value T) (ok bool) {
	_, ok = (*s)[value]
	return
}

func (s *Set[T]) Clear() {
	*s = NewSet[T]()
}
