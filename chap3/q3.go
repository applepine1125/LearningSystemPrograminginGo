package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {

	file, err := os.Create("q3.zip")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("q3.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(writer, strings.NewReader("hello world"))
}
