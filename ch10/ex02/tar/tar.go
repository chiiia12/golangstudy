package tar

import (
	"os"
	"io"
	"log"
	"path/filepath"
	"io/ioutil"
	"archive/tar"
	"../driver"
)

type TarUnArchiver struct {
	inputDir  string
	outputDir string
}

func init() {
	//登録する
	driver.Register("tar", &TarUnArchiver{})
}
func (t *TarUnArchiver) UnArchive(input, output string) {
	unTar(t.inputDir, t.outputDir)
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
