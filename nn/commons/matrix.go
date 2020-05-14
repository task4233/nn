package commons

import (
	"errors"
	"fmt"
	"os"
)

// AreSameMatrixes checks if they are same
func AreSameMatrixes(lArr [][]float64, rArr [][]float64) error {
	if err := CheckMatrixSize(lArr, rArr); err != nil {
		return err
	}

	for ri := 0; ri < len(lArr); ri++ {
		for ci := 0; ci < len(lArr[0]); ci++ {
			if lArr[ri][ci] != rArr[ri][ci] {
				return errors.New("different values")
			}
		}
	}
	return nil
}

// CheckMatrixSize checks structure such as row and colums
func CheckMatrixSize(lArr [][]float64, rArr [][]float64) error {
	status := len(lArr) == len(rArr)
	if !status {
		// for check
		fmt.Fprintf(os.Stderr, "lArr: %v\nrArr: %v\n", lArr, rArr)
		return errors.New("different rows")
	}

	status = len(lArr[0]) == len(rArr[0])
	if !status {
		fmt.Fprintf(os.Stderr, "lArr: %v\nrArr: %v\n", lArr, rArr)
		return errors.New("different columns")
	}

	return nil
}

// CheckTransMatrixSize checks structure of trans matrix for calculating Dot
func CheckTransMatrixSize(lArr [][]float64, rArr [][]float64) error {
	if len(lArr[0]) != len(rArr) {
		fmt.Fprintf(os.Stderr, "lArr: %v\nrArr: %v\n", lArr, rArr)
		return errors.New("different columns")
	}

	return nil
}

// MakeMatrixWithValue makes slice (first arugument) x (second argument)
func MakeMatrixWithValue(row int, column int, value float64) [][]float64 {
	resArr := make([][]float64, row)
	for ri := 0; ri < row; ri++ {
		resArr[ri] = make([]float64, column)
		// init value on make is zero
		if value == 0 {
			continue
		}

		for ci := 0; ci < column; ci++ {
			resArr[ri][ci] = value
		}
	}
	return resArr
}

// Add is the function which add first array and second array
func Add(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	if err := CheckMatrixSize(lArr, rArr); err != nil {
		return nil, fmt.Errorf("failed in add(): %s", err)
	}

	resArr := MakeMatrixWithValue(len(lArr), len(lArr[0]), 0)
	for ri := 0; ri < len(lArr); ri++ {
		for ci := 0; ci < len(lArr[0]); ci++ {
			resArr[ri][ci] = lArr[ri][ci] + rArr[ri][ci]
		}
	}
	return resArr, nil
}

// Dot is the function which multiply first array and second array
func Dot(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	if err := CheckTransMatrixSize(lArr, rArr); err != nil {
		return nil, fmt.Errorf("failed in dot(): %s", err)
	}

	resArr := MakeMatrixWithValue(len(lArr), len(rArr[0]), 0)
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

// ConstMul is the function which multiply the first value and second array
func ConstMul(lVal float64, rArr [][]float64) ([][]float64, error) {
	// implements const multiplication
	resArr := MakeMatrixWithValue(len(rArr), len(rArr[0]), 0)
	for ri := 0; ri < len(resArr); ri++ {
		for ci := 0; ci < len(resArr[0]); ci++ {
			resArr[ri][ci] = rArr[ri][ci] * lVal
		}
	}
	return resArr, nil
}

// Trans returns transverse matrix
func Trans(arr [][]float64) [][]float64 {
	resArr := MakeMatrixWithValue(len(arr[0]), len(arr), 0)
	for ri := 0; ri < len(resArr); ri++ {
		for ci := 0; ci < len(resArr[0]); ci++ {
			resArr[ri][ci] = arr[ci][ri]
		}
	}
	return resArr
}
