package iteration

func Repeat(input string) (output string) {
	for i := 5; i > 0; i-- {
		output = output + input
	}

	return output
}
