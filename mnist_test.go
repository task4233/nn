package nn

import (
	"fmt"
	"testing"
)

func TestMnist(t *testing.T) {
	// set Mnist's data
	mnists := &mnist{
		trainImages: &mnistData{
			filePath:    "./mnist/train-images-idx3-ubyte",
			numOfImages: 0,
			data:        [][][]byte{},
		},
		trainLabels: &mnistLabel{
			filePath:    "./mnist/train-images-idx3-ubyte",
			numOfLabels: 0,
			data:        []int{},
		},
		testImages: &mnistData{
			filePath:    "./mnist/t10k-images-idx3-ubyte",
			numOfImages: 0,
			data:        [][][]byte{},
		},
		testLabels: &mnistLabel{
			filePath:    "./mnist/t10k-labels-idx1-ubyte",
			numOfLabels: 0,
			data:        []int{},
		},
	}

	// load data(true)
	fmt.Println("[TEST] load data(true)")
	if err := mnists.loadDatas(); err != nil {
		t.Errorf("Error: %v\n", err)
	}

	// check size in trainImages
	fmt.Println("[TEST] check size in trainImages")
	if len(mnists.testImages.data) != mnists.testImages.numOfImages {
		t.Errorf("testImages have diff size\n")
		t.Errorf("expected: %d, got: %d", mnists.testImages.numOfImages, len(mnists.testImages.data))
	}

	// check rows in testImages
	fmt.Println("[TEST] check rows in testImages")
	if len(mnists.testImages.data[0]) != mnists.testImages.rows {
		t.Errorf("testImages have diff rows\n")
		t.Errorf("expected: %d, got: %d", mnists.testImages.rows, len(mnists.testImages.data[0]))
	}

	// check columns in testImages
	fmt.Println("[TEST] check rows in testImages")
	if len(mnists.testImages.data[0][0]) != mnists.testImages.columns {
		t.Errorf("testImages have diff columns\n")
		t.Errorf("expected: %d, got: %d", mnists.testImages.columns, len(mnists.testImages.data[0][0]))
	}

	// check size in testLabels
	fmt.Println("[TEST] check size in testLabels")
	if len(mnists.testLabels.data) != mnists.testLabels.numOfLabels {
		t.Errorf("testLabels have diff size\n")
	}

	// check size in trainImages
	fmt.Println("[TEST] check size in trainImages")
	if len(mnists.trainImages.data) != mnists.trainImages.numOfImages {
		t.Errorf("trainImages have diff size\n")
	}

	// check rows in testImages
	fmt.Println("[TEST] check rows in testImages")
	if len(mnists.trainImages.data[0]) != mnists.trainImages.rows {
		t.Errorf("trainImages have diff rows\n")
	}

	// check columns in testImages
	fmt.Println("[TEST] check columns in testImages")
	if len(mnists.trainImages.data[0][0]) != mnists.trainImages.columns {
		t.Errorf("trainImages have diff columns\n")
	}

	// check size in trainLabels
	fmt.Println("[TEST] check size in trainLabels")
	if len(mnists.trainLabels.data) != mnists.trainLabels.numOfLabels {
		t.Errorf("trainLabels have diff size\n")
	}

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp := mnists.trainImages.filePath
	mnists.trainImages.filePath = ""
	if err := mnists.loadDatas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.trainImages.filePath = fp

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp = mnists.trainLabels.filePath
	mnists.trainLabels.filePath = ""
	if err := mnists.loadDatas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.trainLabels.filePath = fp

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp = mnists.testImages.filePath
	mnists.testImages.filePath = ""
	if err := mnists.loadDatas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.testImages.filePath = fp

	// load data(false)
	fmt.Println("[TEST] load data(false)")
	fp = mnists.testLabels.filePath
	mnists.testLabels.filePath = ""
	if err := mnists.loadDatas(); err == nil {
		t.Errorf("Error: %v\n", err)
	}
	mnists.testLabels.filePath = fp
}
