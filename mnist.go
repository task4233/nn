package nn

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type mnist struct {
	train_images *mnist_data
	train_labels *mnist_data
	test_images  *mnist_data
	test_labels  *mnist_data
}

type mnist_data struct {
	filePath    string
	numOfImages int
	data        []byte
}

func (m *mnist) load_datas() error {
	if err := m.train_images.load_data(); err != nil {
		return fmt.Errorf("train_images : %v\n", err)
	}

	if err := m.train_labels.load_data(); err != nil {
		return fmt.Errorf("train_labels : %v\n", err)
	}

	if err := m.test_images.load_data(); err != nil {
		return fmt.Errorf("train_images : %v\n", err)
	}

	if err := m.test_labels.load_data(); err != nil {
		return fmt.Errorf("train_labels : %v\n", err)
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

	// the number of images are in data[4:8]
	m.numOfImages = int(binary.BigEndian.Uint32(data[4:8]))

	m.data = data[8:]

	return nil
}
