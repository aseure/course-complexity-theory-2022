package linear

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	for _, c := range []struct {
		name  string
		stack Stack[int]
	}{
		{"list stack", NewListStack[int]()},
		{"array stack", NewArrayStack[int]()},
	} {
		t.Run(c.name, func(t *testing.T) {
			require.True(t, c.stack.IsEmpty())

			require.True(t, c.stack.Push(0))
			require.False(t, c.stack.IsEmpty())

			require.True(t, c.stack.Push(1))
			require.False(t, c.stack.IsEmpty())

			require.True(t, c.stack.Push(2))
			require.False(t, c.stack.IsEmpty())

			v, ok := c.stack.Pop()
			require.True(t, ok)
			require.Equal(t, 2, v)
			require.False(t, c.stack.IsEmpty())

			v, ok = c.stack.Pop()
			require.True(t, ok)
			require.Equal(t, 1, v)
			require.False(t, c.stack.IsEmpty())

			v, ok = c.stack.Pop()
			require.True(t, ok)
			require.Equal(t, 0, v)
			require.True(t, c.stack.IsEmpty())
		})
	}
}

func BenchmarkStack(b *testing.B) {
	for _, c := range []struct {
		name  string
		stack Stack[int]
	}{
		{"list stack", NewListStack[int]()},
		{"array stack", NewArrayStack[int]()},
	} {
		// TODO Show performances before and after changing the ArrayStack Front/Back implementation
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := 0; i < 10000; i++ {
					c.stack.Push(i)
				}

				for !c.stack.IsEmpty() {
					_, _ = c.stack.Pop()
				}
			}
		})
	}
}
