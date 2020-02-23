package nn

import "math"

// Sigmoid is an activation function.
// f(arr[i]) = 1 / (1 + exp(-arr[i]))
func Sigmoid(arr [][]float64) [][]float64 {
	resArr := makeMatrix(len(arr), len(arr[0]))
	for ri := 0; ri < len(arr); ri++ {
		for ci := 0; ci < len(arr[ri]); ci++ {
			resArr[ri][ci] = 1.0 / (1.0 + math.Exp(-arr[ri][ci]))
		}
	}
	      return resArr, nil
}

// Softmax is an activation function.
// see: https://en.wikipedia.org/wiki/Softmax_function
func Softmax(arr [][]float64) [][]float64 {
	resArr := makeMatrix(len(arr), len(arr[0]))
	for ri := 0; ri < len(arr); ri++ {
		sum := 0.0
		for ci := 0; ci < len(arr[ri]); ci++ {
			sum += math.Exp(arr[ri][ci])
		}

		for ci := 0; ci < len(arr[ri]); ci++ {
			resArr[ri][ci] = math.Exp(arr[ri][ci]) / sum
		}
	}
	return resArr
}
