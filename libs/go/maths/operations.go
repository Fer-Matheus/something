package maths

func Sum(values ...int) (result int) {

	for _, value := range values {
		result += value
	}

	return result
}
