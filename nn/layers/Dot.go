package layers

import (
	"fmt"

	commons "github.com/task4233/nn/nn/commons"
)

// Dot is multiplr larer
type Dot struct {
	lArr [][]float64
	rArr [][]float64
}

// Init is for initialization
func (d *Dot) Init() {
	// init
}

// Forward is for forwarding
func (d *Dot) Forward(lArr [][]float64, rArr [][]float64) ([][]float64, error) {
	d.lArr = lArr
	d.rArr = rArr
	outArr, err := commons.Dot(lArr, rArr)
	if err != nil {
		return nil, fmt.Errorf("failed in Dot.Forward(): %s", err)
	}

	return outArr, nil
}

// Backward is for back propagation
func (d *Dot) Backward(dOutArr [][]float64) ([][]float64, [][]float64, error) {
	dxArr, err := commons.Dot(dOutArr, d.lArr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed in Dot.Backward(): %s", err)
	}
	dyArr, err := commons.Dot(dOutArr, d.rArr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed in Dot.Backward(): %s", err)
	}

	return dxArr, dyArr, nil
}
