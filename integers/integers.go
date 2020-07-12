package integers

func Add(x, y int) (result int) {
	return x + y
}

func Sum(numbers [3]int) (result int) {
	for i := 0; i < 3; i++ {
		result = result + numbers[i]
	}

	return result
}
