package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Michaelpalacce/go-stuff/rle/pkg/encoder"
)

func GetOutFile() (io.WriteCloser, error) {
	filename := "encoded_output.rle"
	// delete if exists
	_, err := os.Stat(filename)
	if err == nil {
		os.Remove(filename)
	}

	// Create a file to write the encoded data
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil, err
	}

	return outputFile, nil
}

func main() {
	outputFile, err := GetOutFile()
	if err != nil {
		panic(err)
	}
	defer outputFile.Close() // Ensure the file is closed when done

	// Create an RleEncoder
	writeStream := encoder.RleEncoder{
		Writer: outputFile,
	}

	// testString := "hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhelloo world!"
	//
	// writeStream.Write(bytes.NewReader([]byte(testString)))

	reader, err := os.Open("test")
	if err != nil {
		panic(err)
	}

	writeStream.Write(reader)
}
