package main

import (
	"fmt"
	"io"
	"os"
)

type File struct {
}

type FileRead interface {
	Read()
	Write()
}

func main() {
	args := os.Args

	var file File
	f, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(file, f)
}

func (*File) Read(b []byte) (int, error) {
	return len(b), nil
}

func (File) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Received string with ", len(b), " Bytes")
	return len(b), nil
}
