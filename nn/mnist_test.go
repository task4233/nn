package nn

import (
	"fmt"
	"testing"
)

func TestMnist(t *testing.T) {
	// set Mnist's data
	mnists := &Mnist{
		TrainImages: &MnistData{
			filePath:    "./Mnist/train-images-idx3-ubyte",
			numOfImages: 0,
			data:        [][][]byte{},
		},
		TrainLabels: &MnistLabel{
			filePath:    "./Mnist/train-images-idx3-ubyte",
			numOfLabels: 0,
			data:        []int{},
		},
		TestImages: &MnistData{
			filePath:    "./Mnist/t10k-images-idx3-ubyte",
			numOfImages: 0,
			data:        [][][]byte{},
		},
		TestLabels: &MnistLabel{
			filePath:    "./Mnist/t10k-labels-idx1-ubyte",
			numOfLabels: 0,
			data:        []int{},
		},
	}

	// load data(true)
	fmt.Println("[TEST] load data(true)")
	if err := mnists.LoadData(); err != nil {
		t.Errorf("Error: %v\n", err)
	}

	// check size in TrainImages
	fmt.Println("[TEST] check size in TrainImages")
	if len(mnists.TestImages.data) != mnists.TestImages.numOfImages {
		t.Errorf("TestImages have diff size\n")
		t.Errorf("expected: %d, got: %d", mnists.TestImages.numOfImages, len(mnists.TestImages.data))
	}

	// check rows in TestImages
	fmt.Println("[TEST] check rows in TestImages")
	if len(mnists.TestImages.data[0]) != mnists.TestImages.rows {
		t.Errorf("TestImages have diff rows\n")
		t.Errorf("expected: %d, got: %d", mnists.TestImages.rows, len(mnists.TestImages.data[0]))
	}

	// check columns in TestImages
	fmt.Println("[TEST] check rows in TestImages")
	if len(mnists.TestImages.data[0][0]) != mnists.TestImages.columns {
		t.Errorf("TestImages have diff columns\n")
		t.Errorf("expected: %d, got: %d", mnists.TestImages.columns, len(mnists.TestImages.data[0][0]))
	}

	// check size in TestLabels
	fmt.Println("[TEST] check size in TestLabels")
	if len(mnists.TestLabels.data) != mnists.TestLabels.numOfLabels {
		t.Errorf("TestLabels have diff size\n")
	}

	// check size in TrainImages
	fmt.Println("[TEST] check size in TrainImages")
	if len(mnists.TrainImages.data) != mnists.TrainImages.numOfImages {
		t.Errorf("TrainImages have diff size\n")
	}

	// check rows in TestImages
	fmt.Println("[TEST] check rows in TestImages")
	if len(mnists.TrainImages.data[0]) != mnists.TrainImages.rows {
		t.Errorf("TrainImages have diff rows\n")
	}

	// check columns in TestImages
	fmt.Println("[TEST] check columns in TestImages")
	if len(mnists.TrainImages.data[0][0]) != mnists.TrainImages.columns {
		t.Errorf("TrainImages have diff columns\n")
	}

	// check size in TrainLabels
	fmt.Println("[TEST] check size in TrainLabels")
	if len(mnists.TrainLabels.data) != mnists.TrainLabels.numOfLabels {
		t.Errorf("TrainLabels have diff size\n")
	}

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp := mnists.TrainImages.filePath
	mnists.TrainImages.filePath = ""
	if err := mnists.LoadData(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.TrainImages.filePath = fp

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp = mnists.TrainLabels.filePath
	mnists.TrainLabels.filePath = ""
	if err := mnists.LoadData(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.TrainLabels.filePath = fp

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp = mnists.TestImages.filePath
	mnists.TestImages.filePath = ""
	if err := mnists.LoadData(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.TestImages.filePath = fp

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp = mnists.TestLabels.filePath
	mnists.TestLabels.filePath = ""
	if err := mnists.LoadData(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.TestLabels.filePath = fp
}
