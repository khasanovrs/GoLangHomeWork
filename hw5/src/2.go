package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("При чтении больших файлов лучше читать файл по кусочкам не загоняя его целиком в память")

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 32*1024) // define your buffer size here.

	for {
		n, err := file.Read(buf)

		if n > 0 {
			fmt.Print(buf[:n]) // your read buffer.
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}
}
