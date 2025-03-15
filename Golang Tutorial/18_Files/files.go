package main

import (
	"bufio"
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
    /* Files in a folder */
	folderInfo, err := dir.ReadDir(0)

	for _, fi := range folderInfo {
		fmt.Println(fi.Name())
	}


	createFile, err := os.Create("example2.txt")
	defer createFile.Close()

	if err != nil {
		panic(err)
	}

	/* Append */
	createFile.WriteString("Hi! Written some content!")
	createFile.WriteString("Hi! Written some content again!")

	/* Append slice of bytes */
	bytes := []byte("Hello!")
	createFile.Write(bytes)

	/* Read and Write (Streaming fashion) */
	sourceFile, err := os.Open("source.txt")
	defer sourceFile.Close()

	if err != nil {
		panic(err)
	}

	/* Write from one file to another */
	destinationFile, err := os.Create("destination.txt")
	defer destinationFile.Close()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destinationFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}

		er := writer.WriteByte(b)
		if er != nil {
			panic(er)
		}

	}

	writer.Flush() 
	fmt.Println("Written...")
	
}