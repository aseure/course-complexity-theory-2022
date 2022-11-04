package sorting

func CountingSort(input []int) []int {
	if len(input) < 2 {
		return input
	}
	max := findMax(input)
	count := make([]int, max+1)
	res := make([]int, len(input))

	for _, v := range input {
		count[v]++
	}

	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}

	for i := len(input) - 1; i >= 0; i-- {
		v := input[i]
		count[v]--
		res[count[v]] = input[i]
	}

	return res
}

func findMax(input []int) int {
	max := input[0]
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	return max
}