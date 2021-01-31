// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/27

package main

import (
	"archive/zip"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Create a buffer to write our archive to.
	//buf := new(bytes.Buffer)
	newZipFile, _ := os.Create("done.zip")
	defer newZipFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(newZipFile)
	defer w.Close()

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
		f.Write([]byte("test line"))
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
