package arrays

func Sum(input []int) int {
	sum := 0
	for _, number := range input {
		sum += number
	}

	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int

	for _, slice := range slicesToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}

	return sums
}
