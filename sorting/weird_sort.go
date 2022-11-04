package sorting

func WeirdSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	return input
}
