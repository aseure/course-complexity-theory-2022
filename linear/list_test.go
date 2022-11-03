package linear

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	for _, c := range []struct {
		name string
		list List[int]
	}{
		{"singly linked list", NewSinglyLinkedList[int]()},
		{"doubly linked list", NewDoublyLinkedList[int]()},
		{"array list", NewArrayList[int]()},
	} {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, 0, c.list.Size())
			require.Len(t, c.list.Values(), 0)
			v, ok := c.list.GetAt(0)
			require.False(t, ok)

			require.True(t, c.list.PushBack(1))
			require.Equal(t, 1, c.list.Size())
			require.Equal(t, []int{1}, c.list.Values())
			v, ok = c.list.GetAt(0)
			require.True(t, ok)
			require.Equal(t, 1, v)

			v, ok = c.list.DeleteAt(0)
			require.Equal(t, 0, c.list.Size())
			require.Len(t, c.list.Values(), 0)
			v, ok = c.list.GetAt(0)
			require.False(t, ok)

			require.True(t, c.list.PushBack(1))
			require.Equal(t, 1, c.list.Size())
			require.Equal(t, []int{1}, c.list.Values())
			v, ok = c.list.GetAt(0)
			require.True(t, ok)
			require.Equal(t, 1, v)

			v, ok = c.list.DeleteAt(0)
			require.Equal(t, 0, c.list.Size())
			require.Len(t, c.list.Values(), 0)
			v, ok = c.list.GetAt(0)
			require.False(t, ok)

			require.True(t, c.list.InsertAt(0, 1))
			require.Equal(t, 1, c.list.Size())
			require.Equal(t, []int{1}, c.list.Values())
			v, ok = c.list.GetAt(0)
			require.True(t, ok)
			require.Equal(t, 1, v)

			require.True(t, c.list.PushBack(2))
			require.Equal(t, 2, c.list.Size())
			require.Equal(t, []int{1, 2}, c.list.Values())
			v, ok = c.list.GetAt(1)
			require.True(t, ok)
			require.Equal(t, 2, v)

			require.True(t, c.list.PushFront(0))
			require.Equal(t, 3, c.list.Size())
			require.Equal(t, []int{0, 1, 2}, c.list.Values())
			v, ok = c.list.GetAt(0)
			require.True(t, ok)
			require.Equal(t, 0, v)

			require.True(t, c.list.InsertAt(1, -1))
			require.Equal(t, 4, c.list.Size())
			require.Equal(t, []int{0, -1, 1, 2}, c.list.Values())
			v, ok = c.list.GetAt(1)
			require.True(t, ok)
			require.Equal(t, -1, v)

			v, ok = c.list.DeleteAt(c.list.Size() - 1)
			require.True(t, ok)
			require.Equal(t, 2, v)
			require.Equal(t, 3, c.list.Size())
			require.Equal(t, []int{0, -1, 1}, c.list.Values())

			v, ok = c.list.DeleteAt(0)
			require.True(t, ok)
			require.Equal(t, 0, v)
			require.Equal(t, 2, c.list.Size())
			require.Equal(t, []int{-1, 1}, c.list.Values())

			v, ok = c.list.DeleteAt(c.list.Size() - 1)
			require.True(t, ok)
			require.Equal(t, 1, v)
			require.Equal(t, 1, c.list.Size())
			require.Equal(t, []int{-1}, c.list.Values())

			v, ok = c.list.DeleteAt(0)
			require.True(t, ok)
			require.Equal(t, -1, v)
			require.Equal(t, 0, c.list.Size())
			require.Len(t, c.list.Values(), 0)
		})
	}
}
