package sorting

func MergeSort(input []int) []int {
	mergeSort(input)
	return input
}

func mergeSort(input []int) {
	if len(input) == 1 {
		return
	}
	middle := len(input) / 2
	mergeSort(input[:middle])
	mergeSort(input[middle:])
	merge(input, middle)
}

func merge(input []int, middle int) {
	var res []int
	i, j := 0, middle

	for i < middle || j < len(input) {
		if i == middle {
			res = append(res, input[j])
			j++
		} else if j == len(input) {
			res = append(res, input[i])
			i++
		} else {
			if input[i] < input[j] {
				res = append(res, input[i])
				i++
			} else {
				res = append(res, input[j])
				j++
			}
		}
	}

	copy(input, res)
}
