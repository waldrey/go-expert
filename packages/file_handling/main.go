package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	//size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Escrevendo dados no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho de %d bytes\n\n", size)
	f.Close()

	// Read files
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Read chunks
	file_chunk, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file_chunk)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	file_chunk.Close()

	// Delete file
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
