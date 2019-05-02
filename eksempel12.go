package main

import (
    "log"
    "os"
)

func main() {
    originalPath := "ww.txt"
    newPath := "www.txt"
    err := os.Rename(originalPath, newPath)
    if err != nil {
        log.Fatal(err)
    }
}