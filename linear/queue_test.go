package linear

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	for _, c := range []struct {
		name  string
		queue Queue[int]
	}{
		{"list queue", NewListQueue[int]()},
	} {
		t.Run(c.name, func(t *testing.T) {
			require.True(t, c.queue.IsEmpty())

			require.True(t, c.queue.Push(0))
			require.False(t, c.queue.IsEmpty())

			require.True(t, c.queue.Push(1))
			require.False(t, c.queue.IsEmpty())

			require.True(t, c.queue.Push(2))
			require.False(t, c.queue.IsEmpty())

			v, ok := c.queue.Pop()
			require.True(t, ok)
			require.Equal(t, 0, v)
			require.False(t, c.queue.IsEmpty())

			v, ok = c.queue.Pop()
			require.True(t, ok)
			require.Equal(t, 1, v)
			require.False(t, c.queue.IsEmpty())

			v, ok = c.queue.Pop()
			require.True(t, ok)
			require.Equal(t, 2, v)
			require.True(t, c.queue.IsEmpty())
		})
	}
}
