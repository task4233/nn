package nn

import (
	"testing"
    "fmt"
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
    fmt.Println("[TEST] load data(true)")
	if err := mnists.load_datas(); err != nil {
		t.Errorf("Error: %v\n", err)
	}

    // check size in train_images
    fmt.Println("[TEST] check size in train_images")
    if len(mnists.test_images.data) != mnists.test_images.numOfImages {
        t.Errorf("test_images have diff size\n")
        t.Errorf("expected: %d, got: %d", mnists.test_images.numOfImages, len(mnists.test_images.data))
    }

    // check rows in test_images
    fmt.Println("[TEST] check rows in test_images")
    if len(mnists.test_images.data[0]) != mnists.test_images.rows {
        t.Errorf("test_images have diff rows\n")
        t.Errorf("expected: %d, got: %d", mnists.test_images.rows, len(mnists.test_images.data[0]))
    }

    // check columns in test_images
    fmt.Println("[TEST] check rows in test_images")
    if len(mnists.test_images.data[0][0]) != mnists.test_images.columns {
        t.Errorf("test_images have diff columns\n")
        t.Errorf("expected: %d, got: %d", mnists.test_images.columns, len(mnists.test_images.data[0][0]))
    }

    // check size in test_labels
    fmt.Println("[TEST] check size in test_labels")
    if len(mnists.test_labels.data) != mnists.test_labels.numOfLabels {
        t.Errorf("test_labels have diff size\n")
    }

    // check size in train_images
    fmt.Println("[TEST] check size in train_images")
    if len(mnists.train_images.data) != mnists.train_images.numOfImages {
        t.Errorf("train_images have diff size\n")
    }

    // check rows in test_images
    fmt.Println("[TEST] check rows in test_images")
    if len(mnists.train_images.data[0]) != mnists.train_images.rows {
        t.Errorf("train_images have diff rows\n")
    }

    // check columns in test_images
    fmt.Println("[TEST] check columns in test_images")
    if len(mnists.train_images.data[0][0]) != mnists.train_images.columns {
        t.Errorf("train_images have diff columns\n")
    }

    // check size in train_labels
    fmt.Println("[TEST] check size in train_labels")
    if len(mnists.train_labels.data) != mnists.train_labels.numOfLabels {
        t.Errorf("train_labels have diff size\n")
    }
    
	// load data(false)
    fmt.Println("[TEST] load data(false)") 
	fp := mnists.train_images.filePath
	mnists.train_images.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.train_images.filePath = fp

	// load data(false)
    fmt.Println("[TEST] load data(false)") 
	fp = mnists.train_labels.filePath
	mnists.train_labels.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.train_labels.filePath = fp

	// load data(false)
    fmt.Println("[TEST] load data(false)") 
	fp = mnists.test_images.filePath
	mnists.test_images.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.test_images.filePath = fp

	// load data(false)
    fmt.Println("[TEST] load data(false)") 
	fp = mnists.test_labels.filePath
	mnists.test_labels.filePath = ""
	if err := mnists.load_datas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.test_labels.filePath = fp
}
