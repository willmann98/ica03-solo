package main

import (
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.fatal(err)
	}

	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

}
