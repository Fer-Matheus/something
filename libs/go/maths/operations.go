package maths

import (
	"fmt"
	"math"
)

func Sum(values ...float64) float64 {
	var result float64
	for _, value := range values {
		result += value
	}

	return result
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, fmt.Errorf("invalid number to calculate square root")
	}
	return math.Sqrt(value), nil
}
