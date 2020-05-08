package commons

import (
	"errors"
	"fmt"
	"os"
)

// checkMatrixSize checks structure such as row and colums
func checkMatrixSize(lArr [][]float64, rArr [][]float64) error {
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

// checkTransMatrixSize checks structure of trans matrix for calculating Dot
func checkTransMatrixSize(lArr [][]float64, rArr [][]float64) error {
	if len(lArr[0]) != len(rArr) {
		fmt.Fprintf(os.Stderr, "lArr: %v\nrArr: %v\n", lArr, rArr)
		return errors.New("different columns")
	}

	return nil
}

// MakeMatrix makes slice (first arugument) x (second argument)
func MakeMatrix(row int, column int) [][]float64 {
	resArr := make([][]float64, row)
	for ri := 0; ri < row; ri++ {
		resArr[ri] = make([]float64, column)
	}
	return resArr
}

// Add is the function which add first array and second array
func Add(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	if err := checkMatrixSize(lArr, rArr); err != nil {
		return nil, fmt.Errorf("failed in add(): %s", err)
	}

	resArr := MakeMatrix(len(lArr), len(lArr[0]))
	for ri := 0; ri < len(lArr); ri++ {
		for ci := 0; ci < len(lArr[0]); ci++ {
			resArr[ri][ci] = lArr[ri][ci] + rArr[ri][ci]
		}
	}
	return resArr, nil
}

// Dot is the function which multiply first array and second array
func Dot(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	if err := checkTransMatrixSize(lArr, rArr); err != nil {
		return nil, fmt.Errorf("failed in dot(): %s", err)
	}

	resArr := MakeMatrix(len(lArr), len(rArr[0]))
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
	resArr := MakeMatrix(len(rArr), len(rArr[0]))
	for ri := 0; ri < len(resArr); ri++ {
		for ci := 0; ci < len(resArr[0]); ci++ {
			resArr[ri][ci] = rArr[ri][ci] * lVal
		}
	}
	return resArr, nil
}

// Trans returns transverse matrix
func Trans(arr [][]float64) [][]float64 {
    resArr := MakeMatrix(len(arr[0]), len(arr))
    for ri:=0; ri<len(resArr); ri++ {
        for ci:=0; ci<len(resArr[0]); ci++ {
            resArr[ri][ci] = arr[ci][ri]
        }
    }
    return resArr
}
