// Copyright 2016 budougumi0617 All Rights Reserved.

// Package tar is based my untar command.
package tar

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/budougumi0617/gotraining/ch10/ex02/unarchive"
)

func unpacking(filename string) (err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	// ReaderはClose()はない。
	reader := tar.NewReader(file)
	fmt.Printf("start !!!!\n")
	var header *tar.Header
	for {
		header, err = reader.Next()
		if err == io.EOF {
			break
		}
		if header.Typeflag == tar.TypeDir {
			log.Printf("%s is directory\n", header.Name)
			os.MkdirAll(header.Name, 0755)
			continue
		}

		if err != nil {
			log.Fatalln(err)
			return
		}

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, reader); err != nil {
			log.Fatalln(err)
			return
		}

		s := "./" + header.Name
		d, _ := filepath.Split(s)

		if _, err = os.Stat(d); err != nil {
			os.MkdirAll(d, 0755)
		}

		if err = ioutil.WriteFile(s, buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Write file ./%s\n", header.Name)
	}
	return nil
}

func init() {
	unarchive.RegisterFormat("tar", "\x75\x73\x74\x61\x72", 257, unpacking)
}
