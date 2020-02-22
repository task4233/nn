package nn

import "fmt"

type neuralNetwork struct {
	middleLayerNum int
	weight         [][][]float64
	bias           [][][]float64
}

// init initalize init values
func (n *neuralNetwork) init() {
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

// forward is forwarding function
func (n *neuralNetwork) forward(dataArr [][]float64) ([][]float64, error) {
	resArr := makeMatrix(len(dataArr), len(dataArr[0]))
	copy(resArr, dataArr)

	for idx := 0; idx < n.middleLayerNum; idx++ {
		tmpArr0, err := dot(resArr, n.weight[idx])
		if err != nil {
			return nil, fmt.Errorf("failed in forward(): %s", err)
		}

		tmpArr1, err := add(tmpArr0, n.bias[idx])
		if err != nil {
			return nil, fmt.Errorf("failed in forward(): %s", err)
		}

		resArr = Sigmoid(tmpArr1)
	}

	return Softmax(resArr), nil
}
