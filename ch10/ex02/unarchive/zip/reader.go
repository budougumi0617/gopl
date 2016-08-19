// Copyright 2016 budougumi0617 All Rights Reserved.

// Package zip is based my unzip command.
package zip

import (
	"archive/zip"
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
	reader, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	return createfile(reader.File)
}

func createfile(files []*zip.File) (err error) {
	var rc io.ReadCloser
	for _, f := range files {
		if f.FileInfo().IsDir() {
			// TODO Cannot unzip subfiles in directory
			fmt.Printf("Cannot unzip subfiles in %s yet\n", f.Name)
		}
		log.Printf("Try name %s\n", f.Name)
		rc, err = f.Open()
		if err != nil {
			log.Fatal(err)
			return
		}
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		if err != nil {
			log.Fatal(err)
			return
		}
		s := "./" + f.Name
		d, _ := filepath.Split(s)
		if _, err = os.Stat(d); err != nil {
			os.MkdirAll(d, 0755)
		}
		if err = ioutil.WriteFile(s, buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Write file ./%s\n", f.Name)
		rc.Close()
	}
	return nil
}

func init() {
	unarchive.RegisterFormat("zip", "\x50\x4b", 0, unpacking)
}
