package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRadix(t *testing.T) {
	for _, c := range []struct {
		n        int
		radix    int
		expected int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{9, 1, 9},
		{12345, 1, 5},
		{12345, 2, 4},
		{12345, 3, 3},
		{12345, 4, 2},
		{12345, 5, 1},
		{123, 4, 0},
	} {
		require.Equal(t, c.expected, getRadix(c.n, c.radix), "n=%d radix=%d", c.n, c.radix)
	}
}

func TestMergeBuckets(t *testing.T) {
	for _, c := range []struct {
		buckets  map[int][]int
		expected []int
	}{
		{
			map[int][]int{
				1: {1, 2, 3},
				2: {4, 5, 6},
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
	} {
		size := 0
		for _, b := range c.buckets {
			size += len(b)
		}
		input := make([]int, size)
		mergeBuckets(input, c.buckets)
		require.Equal(t, c.expected, input, "buckets=%d", c.buckets)
	}
}
