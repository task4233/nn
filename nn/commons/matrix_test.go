package commons

import (
	"fmt"
	"testing"
)

func TestMakeMatrix(t *testing.T) {
	cases := []struct {
		name     string
		input1   int
		input2   int
		expected [][]float64
	}{
		{name: "1x1", input1: 1, input2: 1, expected: [][]float64{{0}}},
		{name: "2x3", input1: 2, input2: 3, expected: [][]float64{{0, 0, 0}, {0, 0, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual := MakeMatrix(c.input1, c.input2)
			if err := CheckMatrixSize(actual, c.expected); err != nil {
				t.Errorf(
					"expected: MakeMatrix(%v, %v) = %v, got %v",
					c.input1, c.input2, c.expected, actual)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		input1   [][]float64
		input2   [][]float64
		expected [][]float64
	}{
		{name: "2x2-2x2", input1: [][]float64{{1, 2}, {2, 3}}, input2: [][]float64{{5, 6}, {1, 2}}, expected: [][]float64{{6, 8}, {3, 5}}},
		{name: "3x3-2x2", input1: [][]float64{{1, 2, 3}, {2, 3, 3}, {3, 1, 2}}, input2: [][]float64{{5, 6}, {1, 2}}, expected: nil},
		{name: "2x2-2x3", input1: [][]float64{{2.3, 3.1}, {3.3, 9.2}}, input2: [][]float64{{56.1, -1.6, -34}, {-1, 0, 0}}, expected: nil}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual, err := Add(c.input1, c.input2)
			if err != nil {
				if c.expected != nil {
					t.Errorf("Error: %v\n", err)
				}
			} else if err := AreSameMatrixes(actual, c.expected); err != nil {
				t.Errorf(
					"expected: Add(%v, %v) = %v, got %v",
					c.input1, c.input2, c.expected, actual)
			}
		})
	}
}

// TestDot is checked on this page
// https://keisan.casio.jp/exec/system/1308269580
func TestDot(t *testing.T) {
	cases := []struct {
		name     string
		input1   [][]float64
		input2   [][]float64
		expected [][]float64
	}{
		{name: "2x2-2x2", input1: [][]float64{{1, 2}, {2, 3}}, input2: [][]float64{{5, 6}, {1, 2}}, expected: [][]float64{{7, 10}, {13, 18}}},
		{name: "3x3-2x2", input1: [][]float64{{1, 2, 3}, {2, 3, 3}, {3, 1, 2}}, input2: [][]float64{{5, 6}, {1, 2}}, expected: nil},
		{name: "2x2-2x3", input1: [][]float64{{2.5, 3.2}, {3.4, -5}}, input2: [][]float64{{56.5, -1.8, -15}, {-1, 0, 0}}, expected: [][]float64{{138.05, -4.5, -37.5}, {197.1, -6.12, -51}}}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual, err := Dot(c.input1, c.input2)
			if err != nil {
				if c.expected != nil {
					t.Errorf("Error: %v\n", err)
				}
			} else if err := AreSameMatrixes(actual, c.expected); err != nil {
				t.Errorf(
					"expected: Dot(%v, %v) = %v, got %v",
					c.input1, c.input2, c.expected, actual)
			}
		})
	}
}

// TestConstMul
func TestConstMul(t *testing.T) {
	cases := []struct {
		name     string
		input1   float64
		input2   [][]float64
		expected [][]float64
	}{
		{name: "val-2x2", input1: 2.0, input2: [][]float64{{2.0, 3.0}, {3.0, 4.0}}, expected: [][]float64{{4.0, 6.0}, {6.0, 8.0}}},
		{name: "0-2x2", input1: 0, input2: [][]float64{{2.0, 3.0}, {3.0, 4.0}}, expected: [][]float64{{0, 0}, {0, 0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual, err := ConstMul(c.input1, c.input2)
			if err != nil {
				if c.expected != nil {
					t.Errorf("Error: %v\n", err)
				}
			} else if err := AreSameMatrixes(actual, c.expected); err != nil {
				t.Errorf(
					"expected: ConstMul(%v, %v) = %v, got %v",
					c.input1, c.input2, c.expected, actual)
			}
		})
	}
}

// TestTrans
func TestTrans(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]float64
		expected [][]float64
	}{
		{name: "2x2", input: [][]float64{{1.0, 2.0}, {3.0, 4.0}}, expected: [][]float64{{1.0, 3.0}, {2.0, 4.0}}},
		{name: "2x3", input: [][]float64{{1.0, 2.0}, {3.0, 4.0}, {5.0, 6.0}}, expected: [][]float64{{1.0, 3.0, 5.0}, {2.0, 4.0, 6.0}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual := Trans(c.input)
			if err := AreSameMatrixes(actual, c.expected); err != nil {
				t.Errorf(
					"expected: Trans(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
