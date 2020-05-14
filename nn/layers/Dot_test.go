package layers

import (
	"fmt"
	"testing"

	commons "github.com/task4233/nn/nn/commons"
)

func TestDotLayer(t *testing.T) {
	dl := &Dot{}
	dl.Init()

	cases := []struct {
		name      string
		input1    [][]float64
		input2    [][]float64
		output    [][]float64
		expected1 [][]float64
		expected2 [][]float64
	}{
		{name: "2x2-2x2", input1: [][]float64{{1, 2}, {2, 3}}, input2: [][]float64{{5, 6}, {1, 2}}, output: [][]float64{{7, 10}, {13, 18}}, expected1: [][]float64{{3, 5}, {3, 5}}, expected2: [][]float64{{6, 8}, {6, 8}}},
		{name: "3x3-2x2", input1: [][]float64{{1, 2, 3}, {2, 3, 3}, {3, 1, 2}}, input2: [][]float64{{5, 6}, {1, 2}}, output: nil, expected1: nil, expected2: nil},
		{name: "2x2-2x3", input1: [][]float64{{2, 3}, {4, -5}}, input2: [][]float64{{1, -1, -4}, {-1, 0, 0}}, output: [][]float64{{-1, -2, -8}, {9, -4, -16}}, expected1: [][]float64{{6, -2}, {6, -2}, {6, -2}}, expected2: [][]float64{{0, -1, -4}, {0, -1, -4}, {0, -1, -4}}},
		{name: "2x2-2x3", input1: [][]float64{{2, 3}, {4, -5}}, input2: [][]float64{{1, -1, -4}, {-1, 0, 0}}, output: [][]float64{{-1, -2, -8}, {9, -4, -16}}, expected1: [][]float64{{6, -2}, {6, -2}, {6, -2}}, expected2: [][]float64{{0, -1, -4}, {0, -1, -4}, {0, -1, -4}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			// test Forward
			output, err := dl.Forward(c.input1, c.input2)
			if err != nil {
				if c.output != nil {
					t.Errorf("Error: %v\n", err)
				}
				return
			} else if err := commons.AreSameMatrixes(output, c.output); err != nil {
				t.Errorf(
					"expected: DotLayer.Forward(%v, %v) = %v, got %v",
					c.input1, c.input2, c.output, output)
			}

			// test Backforward
			onesArr := commons.MakeMatrixWithValue(len(c.output[0]), len(c.output), 1.0)

			actual1, actual2, err := dl.Backward(onesArr)
			if err != nil {
				if c.expected1 != nil {
					t.Errorf("Error: %v\n", err)
				}
			} else if err := commons.AreSameMatrixes(actual1, c.expected1); err != nil {
				t.Errorf(
					"expected: DotLayer.Backward(%v) = %v, %v, got %v, %v",
					onesArr, c.expected1, c.expected2, actual1, actual2)
			} else if err := commons.AreSameMatrixes(actual2, c.expected2); err != nil {
				t.Errorf(
					"expected: DotLayer.Backward(%v) = %v, %v, got %v, %v",
					onesArr, c.expected1, c.expected2, actual1, actual2)
			}
		})
	}
}

// Backward, 転置してからかける必要あるっ書
func TestErrorForBackward(t *testing.T) {
	dl := &Dot{}
	dl.Init()

	lArr := [][]float64{{1, 2, 3}, {2, 3, 3}}
	rArr := [][]float64{{5, 6}, {1, 2}}
	// Foward is OK by checking TestDotLayer
	_, err := dl.Forward(lArr, rArr)
	if err != nil {
		// As it has tested above, it should not be no error.
		t.Errorf("Error: %v\n", err)
	}

	cases := []struct {
		name   string
		output [][]float64
	}{
		{name: "3x3-2x2-1", output: [][]float64{{1}}},
		{name: "3x3-2x2-2", output: [][]float64{{1,1}, {1,1}, {1,1}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			// test Backforward
			onesArr := commons.MakeMatrixWithValue(len(c.output[0]), len(c.output), 1.0)

			_, _, err := dl.Backward(onesArr)
			if err != nil {
				t.Errorf("Error: %v\n", err)
			}
		})
	}
}
