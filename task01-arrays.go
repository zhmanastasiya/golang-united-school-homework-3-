package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	x := len(input)
	for i := 0; i < len(input); i++ {
		sum += input[i]
		result = sum / float32(x)
	}
	return result
}
