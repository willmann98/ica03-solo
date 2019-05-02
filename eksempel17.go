package main

import (
    "os"
    "log"
)

func main() {
    
    file, err := os.Open("ww.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    byteSlice := make([]byte, 16)
    bytesRead, err := file.Read(byteSlice)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Number of bytes read: %d\n", bytesRead)
    log.Printf("Data read: %s\n", byteSlice)
}