package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

// Compress a file called `sparse.img` using `archive/tar`

func main() {
	fmt.Println("Compressing in Go without sparse")
	compress()
	fmt.Println("Compressing in Go with sparse")
	sparseCompress()
}

func compress() {
	err := os.Mkdir("non_sparse", 0777)
	tarfile, err := os.Create("non_sparse/non_sparse.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer tarfile.Close()
	var fileWriter io.WriteCloser = tarfile
	tfw := tar.NewWriter(fileWriter)
	defer tfw.Close()
	file, err := os.Open("sparse.img")
	defer file.Close()
	fileInfo, err := file.Stat()
	header := new(tar.Header)
	header.Name = file.Name()
	header.Size = fileInfo.Size()
	header.Mode = int64(fileInfo.Mode())
	header.ModTime = fileInfo.ModTime()

	err = tfw.WriteHeader(header)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(tfw, file)

	if err != nil {
		log.Fatalln(err)
	}
}

func sparseCompress() {
	err := os.Mkdir("sparse", 0777)
	tarfile, err := os.Create("sparse/sparse.tar")
	if err != nil {
		log.Fatalln(err)
	}
	defer tarfile.Close()
	var fileWriter io.WriteCloser = tarfile
	tfw := tar.NewWriter(fileWriter)
	defer tfw.Close()
	file, err := os.Open("sparse.img")
	defer file.Close()
	fileInfo, err := file.Stat()
	header := new(tar.Header)
	header.Name = file.Name()
	header.Size = fileInfo.Size()
	header.Mode = int64(fileInfo.Mode())
	header.ModTime = fileInfo.ModTime()
	header.Typeflag = tar.TypeGNUSparse

	err = tfw.WriteHeader(header)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(tfw, file)

	if err != nil {
		log.Fatalln(err)
	}
}
