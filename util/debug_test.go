package util

import "testing"

var uid uint32

func BenchmarkUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uid = UID()
	}
}
