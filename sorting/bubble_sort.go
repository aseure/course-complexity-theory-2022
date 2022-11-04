package sorting

func BubbleSort(input []int) []int {
	for hi := len(input) - 1; hi > 0; hi-- {
		for i := 0; i < hi; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	}
	return input
}
