package integers

func Add(x, y int) (result int) {
	return x + y
}

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result = result + number
	}

	return result
}
