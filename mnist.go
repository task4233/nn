package nn

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type mnist struct {
	train_images *mnist_data
	train_labels *mnist_label
	test_images  *mnist_data
	test_labels  *mnist_label
}

type mnist_label struct {
	filePath    string
	numOfLabels int
	data        []int
}

type mnist_data struct {
	filePath    string
	numOfImages int
	rows        int
	columns     int
	// datas HxW(3-dimension)
	data [][][]byte
}

func (m *mnist) load_datas() error {
    fmt.Println("read training images")
	if err := m.train_images.load_data(); err != nil {
		return fmt.Errorf("train_images : %v\n", err)
	}

    fmt.Println("read training labels")
	if err := m.train_labels.load_label(); err != nil {
		return fmt.Errorf("train_labels : %v\n", err)
	}

    fmt.Println("read test images")
	if err := m.test_images.load_data(); err != nil {
		return fmt.Errorf("train_images : %v\n", err)
	}

    fmt.Println("read test labels")
	if err := m.test_labels.load_label(); err != nil {
		return fmt.Errorf("train_labels : %v\n", err)
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
func (m *mnist_label) load_label() error {
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

// load_data loads image data
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
func (m *mnist_data) load_data() error {
	data, err := ioutil.ReadFile(m.filePath)
	if err != nil {
		return err
	}

    // initializing
	// the number of images are in data[4:8]
	m.numOfImages = int(binary.BigEndian.Uint32(data[4:8]))
    m.data = make([][][]byte, m.numOfImages)
    
	m.rows = int(binary.BigEndian.Uint32(data[8:12]))
    for idx:=0; idx<m.numOfImages; idx++ {
        m.data[idx] = make([][]byte, m.rows)
    }
    
	m.columns = int(binary.BigEndian.Uint32(data[12:16]))
    for idx :=0; idx<m.numOfImages; idx++ {
        for ri:=0; ri<m.rows; ri++ {
            m.data[idx][ri] = make([]byte, m.columns)
        }
    }

	sq := m.rows * m.columns

	for idx := 0; idx < m.numOfImages; idx++ {
		for ri := 0; ri < m.rows; ri++ {
			// TODO
			m.data[idx][ri] = data[8+idx*sq : 8+(idx+1)*sq]
		}
	}
	return nil
}

