package sorting

func QuickSort(input []int) []int {
	quickSort(input, 0, len(input)-1)
	return input
}

func quickSort(input []int, low, high int) {
	if low < high {
		pivot := partition(input, low, high)
		quickSort(input, low, pivot-1)
		quickSort(input, pivot+1, high)
	}
}

func partition(input []int, low, high int) int {
	pivot := input[high]
	i := low - 1
	for j := low; j < high; j++ {
		if input[j] < pivot {
			i++
			input[i], input[j] = input[j], input[i]
		}
	}
	input[i+1], input[high] = input[high], input[i+1]
	return i + 1
}
