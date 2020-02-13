package mnist

import (
	"errors"
)

func checkMatrixSize(lArr [][]float64, rArr [][]float64) error {
	status := len(lArr) == len(rArr)
	if !status {
		return errors.New("Different rows")
	}

	status = len(lArr[0]) == len(rArr[0])
	if !status {
		return errors.New("Different columns")
	}

	return nil
}

func checkTransMatrixSize(lArr [][]float64, rArr [][]float64) error {
	if len(lArr[0]) != len(rArr) {
		return errors.New("Different columns")
	}

	return nil
}

// makeMatrix makes slice (first arugument) x (second argument)
func makeMatrix(row int, column int) [][]float64 {
	resArr := make([][]float64, row)
	for ri := 0; ri < row; ri++ {
		resArr[ri] = make([]float64, column)
	}
	return resArr
}

// add is the function which add first array and second array
func add(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	if err := checkMatrixSize(lArr, rArr); err != nil {
		return nil, err
	}

	resArr := makeMatrix(len(lArr), len(lArr[0]))
	for ri := 0; ri < len(lArr); ri++ {
		for ci := 0; ci < len(lArr[0]); ci++ {
			resArr[ri][ci] = lArr[ri][ci] + rArr[ri][ci]
		}
	}
	return resArr, nil
}

// dot is the function which multiply first array and second array
func dot(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	if err := checkTransMatrixSize(lArr, rArr); err != nil {
		return nil, err
	}

	resArr := makeMatrix(len(lArr), len(rArr[0]))
	for ri := 0; ri < len(resArr); ri++ {
		for ci := 0; ci < len(resArr[0]); ci++ {
			sum := 0.0
			for rj := 0; rj < len(lArr[0]); rj++ {
				sum += lArr[ri][rj] * rArr[rj][ci]
			}
			resArr[ri][ci] = sum
		}
	}
	return resArr, nil
}
