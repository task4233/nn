package nn

import (
	"testing"
)

func TestMnist(t *testing.T) {
	// set Mnist's data
	mnists := &mnist{
		train_images: &mnist_data{
			filePath:    "./mnist/train-images-idx3-ubyte",
			numOfImages: 0,
			data:        [][][]byte{},
		},
		train_labels: &mnist_label{
			filePath:    "./mnist/train-images-idx3-ubyte",
			numOfLabels: 0,
			data:        []int{},
		},
		test_images: &mnist_data{
			filePath:    "./mnist/t10k-images-idx3-ubyte",
			numOfImages: 0,
			data:        [][][]byte{},
		},
		test_labels: &mnist_label{
			filePath:    "./mnist/t10k-labels-idx1-ubyte",
			numOfLabels: 0,
			data:        []int{},
		},
	}

	// load data(true)
	if err := mnists.load_datas(); err != nil {
		t.Errorf("Error: %v\n", err)
	}

	// load data(false)
	fp := mnists.train_images.filePath
	mnists.train_images.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.train_images.filePath = fp

	// load data(false)
	fp = mnists.train_labels.filePath
	mnists.train_labels.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.train_labels.filePath = fp

	// load data(false)
	fp = mnists.test_images.filePath
	mnists.test_images.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.test_images.filePath = fp

	// load data(false)
	fp = mnists.test_labels.filePath
	mnists.test_labels.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.test_labels.filePath = fp

}
