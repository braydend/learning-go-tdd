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
	var sums []int

	for _, collection := range collectionsToSum {
		if len(collection) == 0 {
			sums = append(sums, 0)
		} else {
			tail := collection[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
