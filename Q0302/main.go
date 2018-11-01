package main

import (
	"crypto/rand"
	"os"
)

func main() {
	file, err := os.Create("rand.bin")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	_, err = rand.Reader.Read(buffer)
	if err != nil {
		panic(err)
	}
	file.Write(buffer)
}
