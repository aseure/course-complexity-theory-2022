package tree

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap[int]()
	require.Equal(t, 0, h.Size())

	values := []int{
		2,
		19,
		100,
		3,
		36,
		7,
		17,
		25,
		1,
	}
	for _, v := range values {
		h.Insert(v)
	}
	require.Equal(t, len(values), h.Size())

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	for _, v := range values {
		max, ok := h.RemoveMax()
		require.True(t, ok)
		require.Equal(t, v, max)
	}
}
