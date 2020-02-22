package nn

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type mnist struct {
	trainImages *mnistData
	trainLabels *mnistLabel
	testImages  *mnistData
	testLabels  *mnistLabel
}

type mnistLabel struct {
	filePath    string
	numOfLabels int
	data        []int
}

type mnistData struct {
	filePath    string
	numOfImages int
	rows        int
	columns     int
	// datas HxW(3-dimension)
	data [][][]byte
}

func (m *mnist) loadDatas() error {
	fmt.Println("read training images")
	if err := m.trainImages.loadData(); err != nil {
		return fmt.Errorf("train images : %v", err)
	}

	fmt.Println("read training labels")
	if err := m.trainLabels.loadLabel(); err != nil {
		return fmt.Errorf("train labels : %v", err)
	}

	fmt.Println("read test images")
	if err := m.testImages.loadData(); err != nil {
		return fmt.Errorf("train images : %v", err)
	}

	fmt.Println("read test labels")
	if err := m.testLabels.loadLabel(); err != nil {
		return fmt.Errorf("train labels : %v", err)
	}
	return nil
}

// [offset] [type]          [value]          [description]
// 0000     32 bit integer  0x00000801(2049) magic number (MSB first)
// 0004     32 bit integer  60000            number of items
// 0008     unsigned byte   ??               label
// 0009     unsigned byte   ??               label
// ........
// xxxx     unsigned byte   ??               label
// The labels values are 0 to 9.
func (m *mnistLabel) loadLabel() error {
	data, err := ioutil.ReadFile(m.filePath)
	if err != nil {
		return err
	}

	// the number of labels are in data[4:8]
	m.numOfLabels = int(binary.BigEndian.Uint32(data[4:8]))
	m.data = make([]int, m.numOfLabels)

	for idx := 0; idx < m.numOfLabels; idx++ {
		m.data[idx] = int(data[8+idx])
	}
	return nil
}

// loadData loads image data
// TRAINING SET IMAGE FILE (train-images-idx3-ubyte):
// [offset] [type]          [value]          [description]
// 0000     32 bit integer  0x00000803(2051) magic number
// 0004     32 bit integer  60000            number of images
// 0008     32 bit integer  28               number of rows
// 0012     32 bit integer  28               number of columns
// 0016     unsigned byte   ??               pixel
// 0017     unsigned byte   ??               pixel
// ........
// xxxx     unsigned byte   ??               pixel
func (m *mnistData) loadData() error {
	data, err := ioutil.ReadFile(m.filePath)
	if err != nil {
		return err
	}

	// initializing
	// the number of images are in data[4:8]
	m.numOfImages = int(binary.BigEndian.Uint32(data[4:8]))
	m.data = make([][][]byte, m.numOfImages)

	m.rows = int(binary.BigEndian.Uint32(data[8:12]))
	for idx := 0; idx < m.numOfImages; idx++ {
		m.data[idx] = make([][]byte, m.rows)
	}

	m.columns = int(binary.BigEndian.Uint32(data[12:16]))
	for idx := 0; idx < m.numOfImages; idx++ {
		for ri := 0; ri < m.rows; ri++ {
			m.data[idx][ri] = make([]byte, m.columns)
		}
	}

	sq := m.rows * m.columns

	for idx := 0; idx < m.numOfImages; idx++ {
		// assign data each rows
		// m.data[idx][rows][columns]
		// m.data[idx=0][row=0] = data[8 + idx=0*sq + 0 * columns:8 + 1 * columns]
		// m.data[idx=0][row=1] = data[8 + idx=0*sq + 1 * columns:8 + 2 * columns]
		// ...
		// m.data[idx=1][row=0] = data[8 + idx=1*sq + 0 * columns:8 + 1 * columns]
		for ri := 0; ri < m.rows; ri++ {
			lb := 8 + idx*sq + ri*m.columns
			ub := 8 + idx*sq + (ri+1)*m.columns
			m.data[idx][ri] = data[lb:ub]
		}
	}
	return nil
}