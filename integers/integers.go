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

func SumAllTails(collectionsToSum ...[]int) []int {
	numberOfCollections := len(collectionsToSum)
	sums := make([]int, numberOfCollections)

	for i, collection := range collectionsToSum {
		numbers := collection[1:]
		sums[i] = Sum(numbers)
	}

	return sums
}
