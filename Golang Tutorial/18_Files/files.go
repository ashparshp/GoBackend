package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("Filename: ", fileInfo.Name())
	fmt.Println("isFolder: ", fileInfo.IsDir())
	fmt.Println("sizeOf: ", fileInfo.Size(), "Byte")
	fmt.Println("File Permission: ", fileInfo.Mode())
	fmt.Println("Last Modified: ", fileInfo.ModTime())


	/* Read */
	buf := make([]byte, fileInfo.Size())

	datalen, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data Length: ", datalen, string(buf))

	for i := range buf {
		fmt.Println("Data: ", datalen, string(buf[i]))
	}

	/* Simple Read */
	read, err := os.ReadFile("example.txt")
	fmt.Println(string(read))

	/* Read Folders */
	dir, err := os.Open(".")
	defer dir.Close()

	if err != nil {
		panic(err)
	}

}