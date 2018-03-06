package main

import (
	"log"
	"archive/zip"
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"archive/tar"
	"image"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := unZip("./sample.zip", "./out")
	err = unTar("./sample.tar", "./out")
	if err != nil {
		log.Println("unTar return err", err)
	}

	sql.Open("mysql","dbname")
}
func unZip(input, output string) error {
	r, err := zip.OpenReader(input)
	if err != nil {
		log.Println(err)
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		if f.FileInfo().IsDir() {
			path := filepath.Join(output, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize64)
			_, err := io.ReadFull(rc, buf)
			if err != nil {
				return err
			}
			path := filepath.Join(output, f.Name)
			if err = ioutil.WriteFile(path, buf, f.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func unTar(in, out string) error {
	file, _ := os.Open(in)
	tarReader := tar.NewReader(file)
	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		log.Println("tarHeader.Name", tarHeader.Name)
		f, _ := os.Open(tarHeader.Name)
		info, _ := f.Stat()
		log.Println("fileName is ", info.Name())
		if info.IsDir() {
			path := filepath.Join(out, info.Name())
			log.Println("path is dir: ", path)
			os.MkdirAll(path, info.Mode())
		} else {
			buf := []byte{}
			log.Println("this is file ", f.Name())
			_, err := io.ReadFull(file, buf)
			if err != nil {
				log.Println("73 is :", err)
				return err
			}
			path := filepath.Join(out, f.Name())
			log.Println("path is file: ", path)
			if err = ioutil.WriteFile(path, buf, info.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}
