package commons

import (
	"fmt"
	"testing"
)

// TestSigmoid is checked on python3
// accuracy is the 16-th decimal
func TestSigmoid(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]float64
		expected [][]float64
	}{
		{name: "1x1", input: [][]float64{{0}}, expected: [][]float64{{0.5}}},
		{name: "2x3", input: [][]float64{{1.1, 0, 2.2}, {-4, -1.5, 125}}, expected: [][]float64{{0.7502601055951177, 0.5, 0.9002495108803148}, {0.01798620996209156, 0.18242552380635632, 1}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Sigmoid(c.input)
			if err := AreSameMatrixes(actual, c.expected); err != nil {
				t.Errorf(
					"expected: Sigmoid(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}

func TestSoftmax(t *testing.T) {
	cases := []struct {
		name     string
		input    [][]float64
		expected [][]float64
	}{
		{name: "zero division", input: [][]float64{{0}}, expected: [][]float64{{1}}},
		{name: "1x1", input: [][]float64{{1}}, expected: [][]float64{{1}}},
		{name: "2x3", input: [][]float64{{1.1, 0, 2.2}, {-4, -1.5, 125}}, expected: [][]float64{{0.23057215679279944, 0.07675080370222266, 0.6926770395049778}, {9.462629465836377e-57, 1.152784263199264e-55, 1}}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fmt.Printf("[TEST] %s begins\n", c.name)
			actual := Softmax(c.input)
			if err := AreSameMatrixes(actual, c.expected); err != nil {
				t.Errorf(
					"expected: Softmax(%v) = %v, got %v",
					c.input, c.expected, actual)
			}
		})
	}
}
