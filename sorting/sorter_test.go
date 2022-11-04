package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const size = 100_000

func TestSorter(t *testing.T) {
	shuffled, sorted, reversed := makeInput(size)
	require.NotEqual(t, sorted, shuffled)

	for _, c := range []struct {
		name   string
		sorter Sorter
	}{
		{"Bubble Sort", BubbleSort},
		{"Weird Sort", WeirdSort},
		{"Merge Sort", MergeSort},
		{"Quick Sort", QuickSort},
		{"Radix Sort", RadixSort},
		{"Counting Sort", CountingSort},
	} {
		tmp := make([]int, len(shuffled))
		copy(tmp, shuffled)
		start := time.Now()
		assert.Equal(t, sorted, c.sorter(tmp), c.name)
		fmt.Printf("%s (shuffled) took %s\n", c.name, time.Since(start))

		copy(tmp, sorted)
		start = time.Now()
		assert.Equal(t, sorted, c.sorter(tmp), c.name)
		fmt.Printf("%s (sorted) took %s\n", c.name, time.Since(start))

		copy(tmp, reversed)
		start = time.Now()
		assert.Equal(t, sorted, c.sorter(tmp), c.name)
		fmt.Printf("%s (reversed) took %s\n", c.name, time.Since(start))

		fmt.Println()
	}
}

func makeInput(size int) (shuffled, sorted, reversed []int) {
	shuffled = make([]int, size)
	sorted = make([]int, size)
	reversed = make([]int, size)
	for i := range shuffled {
		shuffled[i] = i
		sorted[i] = i
		reversed[size-1-i] = i
	}
	rand.Seed(int64(size))
	rand.Shuffle(len(shuffled), func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
	return
}
