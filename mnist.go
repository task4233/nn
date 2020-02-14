package main

import (
	"fmt"
	"io/ioutil"
	"os"
    "encoding/binary"
)

// load_images loads image data
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
func load_images(fp string) ([]byte, error) {
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return nil, err
	}

    // the number of images are in data[4:8]
	numOfImages := binary.BigEndian.Uint32(data[4:8])
    fmt.Println(numOfImages)

	resData := make([]byte, numOfImages)
	resData = data[8:]

	return resData, nil
}

func main() {
	fps := []string{"./mnist/t10k-images-idx3-ubyte", "./mnist/t10k-labels-idx1-ubyte", "./mnist/train-images-idx3-ubyte", "./mnist/train-labels-idx1-ubyte"}
	for _, fp := range fps {
		fmt.Printf("%s : ", fp)
		resData, err := load_images(fp)
        if err != nil {
			fmt.Errorf("Error: %v\n", err)
			os.Exit(1)
		}

        fmt.Println(resData)
        
	}

}
