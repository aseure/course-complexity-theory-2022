package sorting

func RadixSort(input []int) []int {
	radix := 1
	for {
		buckets := makeRadixBuckets()
		for _, n := range input {
			bucketID := getRadix(n, radix)
			buckets[bucketID] = append(buckets[bucketID], n)
		}
		radix++
		mergeBuckets(input, buckets)
		if len(buckets[0]) == len(input) {
			return input
		}
	}
}

func makeRadixBuckets() map[int][]int {
	return map[int][]int{
		0: nil,
		1: nil,
		2: nil,
		3: nil,
		4: nil,
		5: nil,
		6: nil,
		7: nil,
		8: nil,
		9: nil,
	}
}

func getRadix(n, radix int) int {
	for i := 1; i < radix; i++ {
		n /= 10
	}
	return n % 10
}

func mergeBuckets(dst []int, buckets map[int][]int) {
	start := 0
	for i := 0; i < 10; i++ {
		end := start + len(buckets[i])
		copy(dst[start:end], buckets[i])
		start = end
	}
}
