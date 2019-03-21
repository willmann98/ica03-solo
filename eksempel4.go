package main

import (
	"fmt"
	"log"
	"os"
)

const fil = "C:/Users/willi/go/src/github.com/Willmann98/is105-ica03/files/pg100.txt"

var (
	fileInfo os.FileInfo
	err      error
)

func main() {

	fileInfo, err = os.Stat("fil")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.sys())

}
