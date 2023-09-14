package util

type Stack[T any] []T

func NewStack[T any](initialCapacity int) Stack[T] {
	return make(Stack[T], 0, initialCapacity)
}

func (s *Stack[T]) Push(value ...T) *T {
	*s = append(*s, value...)
	return s.TopRef()
}

func (s *Stack[T]) top() int {
	return len(*s) - 1
}

func (s *Stack[T]) Top() (result T, ok bool) {
	if top := s.top(); top >= 0 {
		result = (*s)[top]
		ok = true
	}
	return
}

func (s *Stack[T]) TopRef() (result *T) {
	if top := s.top(); top >= 0 {
		result = &(*s)[top]
	}
	return
}

func (s *Stack[T]) Pop() (value *T) {
	if value = s.TopRef(); value != nil {
		*s = (*s)[:len(*s)-1]
	}
	return
}

func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}
