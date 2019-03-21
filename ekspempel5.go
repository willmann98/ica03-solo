package main

import (
	"io"
	"log"
	"os"
)

func main() {

	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("copied %d bytes.", bytesWritten)

	err = newFile.sync()
	if err != nil {
		log.Fatal(err)
	}

}
