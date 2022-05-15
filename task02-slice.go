package homework

func reverse(input []int64) (result []int64) {
	input2 := make([]int64, len(input))
	copy(input2, input)
	for i := len(input2)/2 - 1; i >= 0; i-- {
		opp := len(input2) - 1 - i
		input2[i], input2[opp] = input2[opp], input2[i]
	}
	return input2
}
