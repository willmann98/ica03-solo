package main

import (
	"log"
	"os"
)

const fil = "C:/Users/willi/go/src/github.com/Willmann98/is105-ica03/fileio/FilTest.txt"

func main() {

	file, err := os.Open(fil)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.OpenFile(fil, os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	
}
