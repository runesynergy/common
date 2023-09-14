package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var test_result string

func TestStack(t *testing.T) {
	var stack Stack[string]
	require.Equal(t, 0, len(stack))
	stack.Push("Hello, world!")
	require.Equal(t, 1, len(stack))
	value := stack.Pop()
	require.Equal(t, "Hello, world!", value)
	require.Equal(t, 0, len(stack))
}

func BenchmarkStackTop(b *testing.B) {
	b.ReportAllocs()
	var stack Stack[string]
	stack.Push("Hello, world!")
	for i := 0; i < b.N; i++ {
		test_result, _ = stack.Top()
	}
}
func BenchmarkStackTopRef(b *testing.B) {
	b.ReportAllocs()
	var stack Stack[string]
	stack.Push("Hello, world!")
	for i := 0; i < b.N; i++ {
		ref := stack.TopRef()
		test_result = *ref
	}
}

func BenchmarkStackPushPop(b *testing.B) {
	b.ReportAllocs()
	var stack Stack[string]
	for i := 0; i < b.N; i++ {
		stack.Push("Hello, world!")
		test_result = *stack.Pop()
	}
}
