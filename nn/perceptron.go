package nn

import (
    "fmt"
    commons "github.com/task4233/nn/nn/commons"
)

// NeuralNetwork has the information for building neural network
type NeuralNetwork struct {
	middleLayerNum int
	weight         [][][]float64
	bias           [][][]float64
}

// Init initialize init values
func (n *NeuralNetwork) Init() {
	n.middleLayerNum = 2
	n.weight = [][][]float64{
		{{0.1, 0.1, 0.1}, {0.1, 0.1, 0.1}},
		{{0.2, 0.2}, {0.2, 0.2}, {0.3, 0.3}},
	}
	n.bias = [][][]float64{
		{{0.2, 0.2, 0.2}},
		{{0.1, 0.1}},
	}
}

// Forward is forwarding function
func (n *NeuralNetwork) Forward(dataArr [][]float64) ([][]float64, error) {
	resArr := commons.MakeMatrix(len(dataArr), len(dataArr[0]))
	copy(resArr, dataArr)

	for idx := 0; idx < n.middleLayerNum; idx++ {
		tmpArr0, err := commons.Dot(resArr, n.weight[idx])
		if err != nil {
			return nil, fmt.Errorf("failed in forward(): %s", err)
		}

		tmpArr1, err := commons.Add(tmpArr0, n.bias[idx])
		if err != nil {
			return nil, fmt.Errorf("failed in forward(): %s", err)
		}

		resArr = commons.Sigmoid(tmpArr1)
	}

	return commons.Softmax(resArr), nil
}

// Backward is back forwarding function
func (n *NeuralNetwork) Backward(dataArr [][]float64) ([][]float64, error) {
    resArr := commons.MakeMatrix(len(dataArr), len(dataArr[0]))

    // dummy value for return
    return resArr, nil
}
