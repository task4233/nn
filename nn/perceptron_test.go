package nn

import (
	"fmt"
	"testing"
    commons "github.com/task4233/nn/nn/commons"
)

func TestForward(t *testing.T) {
	nn := &NeuralNetwork{}
	nn.Init()

	cases := []struct {
		name     string
		input    [][]float64
		expected [][]float64
	}{
		{name: "test forwarding", input: [][]float64{{1.0, 0.5}}, expected: [][]float64{{0, 0}}},
		{name: "diff rows", input: [][]float64{{1.0, 0.5}, {2.0, 0.5}}, expected: nil},
		{name: "diff columns", input: [][]float64{{1.0, 0.5, 1.5}}, expected: nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual, err := nn.Forward(c.input)
			if err != nil {
				if c.expected != nil {
					t.Errorf("Error: %v\n", err)
				}
			} else if err := commons.CheckMatrixSize(actual, c.expected); err != nil {
				t.Errorf("Error: %v\n", err)
			}
		})
	}

}
