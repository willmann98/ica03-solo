package main

import (
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("ww.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}
