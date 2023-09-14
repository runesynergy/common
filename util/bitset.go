package util

import "golang.org/x/exp/constraints"

type Bitset[F constraints.Unsigned] uint64

func (s *Bitset[F]) Has(f F) bool {
	return F(*s)&f == f
}

func (s *Bitset[F]) Add(f F) {
	*s |= Bitset[F](f)
}

func (s *Bitset[F]) Remove(f F) {
	*s &^= Bitset[F](f)
}
