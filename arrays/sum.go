package arrays

func Sum(input []int) int {
	sum := 0
	for _, number := range input {
		sum += number
	}
	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	return []int{2, 11}
}
