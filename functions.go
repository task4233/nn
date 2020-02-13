package mnist

import (
	"errors"
	"math"
)

// Sigmoid is an activation function
// 1 / (1 + exp(-arr[i]))
func Sigmoid(arr [][]float64) [][]float64 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = 1.0 / (1.0 + math.Exp(-arr[i][j]))
		}
	}
	return arr
}

// Softmax is an activation function
func Softmax(arr [][]float64) ([][]float64, error) {
	for i := 0; i < len(arr); i++ {
		sum := 0.0
		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}

		if sum == 0 {
			return nil, errors.New("zero division")
		}

		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] /= sum
		}
	}
	return arr, nil
}
