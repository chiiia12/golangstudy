package zip

import (
	"log"
	"path/filepath"
	"os"
	"io"
	"io/ioutil"
	"archive/zip"
	"../driver"
	"fmt"
)

type ZipUnArchiver struct {
}

func init() {
	//登録する
	driver.Register("zip", &ZipUnArchiver{})
}

func (z *ZipUnArchiver) UnArchive(input, output string) {
	unZip(input, output)
}

func unZip(input, output string) error {
	r, err := zip.OpenReader(input)
	if err != nil {
		log.Println("31:", err)
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		log.Println("file is ", f)
		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("f.Open return error.%v", err)
		}
		defer rc.Close()
		if f.FileInfo().IsDir() {
			path := filepath.Join(output, f.Name)
			output = path
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize64)
			_, err := io.ReadFull(rc, buf)
			if err != nil {
				return fmt.Errorf("io.ReadFull return error.%v", err)
			}
			path := filepath.Join(output, f.Name)
			output = path
			if err = ioutil.WriteFile(path, buf, f.Mode()); err != nil {
				return fmt.Errorf("ioutil.WriteFile return error.%v", err)
			}
		}
	}
	return nil
}
